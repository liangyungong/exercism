(ns bird-watcher)

(def last-week
  [0 2 5 3 7 8 4])

(defn today [birds]
  (last birds))

(defn inc-bird [birds]
  let [today_count (today birds)
       before_today (drop-last birds)]
  (conj before_today (+ today_count 1)))

(defn day-without-birds? [birds]
  (any (map (fn [c] (= c 0) birds))))

(defn n-days-count [birds n])

(defn busy-days [birds])

(defn odd-week? [birds])
