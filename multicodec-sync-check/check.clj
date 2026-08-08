#!/usr/bin/env bb
(ns check
  (:require [babashka.fs :as fs]
            [clojure.data.csv :as csv]
            [clojure.java.io :as io]
            clojure.string
            clojure.pprint))

(def multiaddr-dir (.getParent (io/file (.getParent (io/file *file*)))))

(def multicodec-csv-file
  (cond
    (fs/exists? (io/file multiaddr-dir "../multicodec")) (io/file multiaddr-dir "../multicodec/table.csv")
    (fs/exists? (io/file multiaddr-dir "./multicodec")) (io/file multiaddr-dir "./multicodec/table.csv")
    :else
    (do
      (println "Can't find multicodec repo")
      (System/exit 1))))

(defn parseHex [h]
  (-> h
      (clojure.string/split #"0x")
      second
      (Integer/parseInt 16)))

(defn parse-csv [reader code-parser]
  (let [data (csv/read-csv reader)
        headers (map (comp keyword clojure.string/trim) (first data))
        body (map #(map clojure.string/trim %) (rest data))]
    (doall
     (->>
      body
      (map (partial zipmap headers))
      (map #(update % :code code-parser))))))

(def multicodec-contents
  (with-open [reader (io/reader multicodec-csv-file)]
    (parse-csv reader parseHex)))

(def multiaddr-contents
  (with-open [reader (io/reader (io/file multiaddr-dir "protocols.csv"))]
    (parse-csv reader #(Integer/parseInt %))))

(defn to-codec-map [table] (reduce #(assoc %1 (:code %2) %2) {} table))

(def missing-multicodecs
  (let [multicodec-map (to-codec-map multicodec-contents)]
    (reduce
     (fn [acc multiaddr]
       (if-not (contains? multicodec-map (:code multiaddr))
         (conj acc multiaddr)
         acc))
     [] multiaddr-contents)))


(when (> (count missing-multicodecs) 0)
  (println "Some protocols in the multiaddr table are not registered with multicodecs:")
  (clojure.pprint/print-table missing-multicodecs)
  (System/exit 1))