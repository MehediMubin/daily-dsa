# 🗳️ Boyer-Moore Voting Algorithm

The **Boyer-Moore Voting Algorithm** is an efficient way to find the **majority element** in an array — an element that appears **more than ⌊n / 2⌋ times**.

> 💡 It guarantees the correct result only if a majority element **exists**.

---

## 🧠 Idea

The algorithm works by **pairing different elements** and canceling them out. If a majority element exists, it will be the one left at the end.

---

## 👨‍💻 Steps

1. Initialize `count = 0`, and let `candidate` be `None`.
2. Iterate over the array:
   -  If `count == 0`, assign the current element as the `candidate`.
   -  If the current element is equal to `candidate`, increment `count`.
   -  Otherwise, decrement `count`.
3. After one pass, `candidate` will hold the majority element.

---

## ⏱ Complexity

-  **Time:** O(n)
-  **Space:** O(1)

---

## ✅ Example

```go
func majorityElement(nums []int) int {
    cnt := 0
    candidate := 0
    for _, num := range(nums) {
        if cnt == 0 {
            candidate = num
        }
        if (candidate == num) {
            cnt++
        } else {
            cnt--
        }
    }
    return candidate
}
```
