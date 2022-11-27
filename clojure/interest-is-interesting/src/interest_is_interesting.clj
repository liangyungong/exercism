(ns interest-is-interesting)

(defn interest-rate
  "return the interest rate based balance"
  [balance]
  (cond
    (< balance 0M) -3.213
    (< balance 1000M) 0.5
    (< balance 5000M) 1.621
    :else 2.475))

(defn annual-balance-update
  "=> balance + interest"
  [balance]
  (let [bigdec-interest (bigdec (interest-rate balance))
        sign (if (pos? balance) + -)]
    (sign (* balance (sign (/ bigdec-interest 100M) 1M)))))

(defn amount-to-donate
  "=> 2x tax-free"
  [balance tax-free-percentage]
  (let [b-tax-free-percentage (bigdec tax-free-percentage)]
    (let [result (int (/ (* balance b-tax-free-percentage 2) 100))]
      (if (pos? result)
        result
        0))))
