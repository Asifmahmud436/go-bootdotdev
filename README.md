# GO Progress

---

### ১. Go data types এবং variable declaration

- Go তে কি কি primitive data types আছে?  
- variable declare করার ৩ ধরনের syntax কী কী?  
- zero value কী? কোন type এর zero value কী?  

**Practice:**  
- একটা struct তৈরি করো যেটার মধ্যে name, age, isActive ফিল্ড থাকবে।  
- zero value দিয়ে একটি variable declare করো, সেটার ভ্যালু print করো।  

---

### ২. Control flow (if, for, switch)

- Go তে for loop এর ৩ ধরনের ব্যবহার কী কী?  
- switch এর syntax এবং fallthrough কী?  
- defer কীভাবে কাজ করে?  

**Practice:**  
- একটা ফাংশন লিখো যা ১ থেকে ১০ পর্যন্ত even সংখ্যা print করবে।  
- একটা switch লিখো যেটা string input অনুযায়ী “Good Morning”, “Good Afternoon”, “Good Evening” print করবে।  

---

### ৩. Functions এবং Methods

- Go তে function declare করার syntax কী?  
- multiple return values কিভাবে handle করা হয়?  
- method receiver কী? value receiver আর pointer receiver এর পার্থক্য কী?  

**Practice:**  
- একটা function লিখো যা দুইটা int নেবে এবং তাদের যোগফল আর গুণফল রিটার্ন করবে।  
- struct তৈরি করে তার উপর method add করো যেটা struct এর একটি ফিল্ড modify করবে।  

---

### ৪. Pointers

- pointer কি? Go তে pointer declare কিভাবে করে?  
- pointer dereference এবং pointer arithmetic Go তে কেমন?  
- pointer আর value receiver method এর পার্থক্য কী?  

**Practice:**  
- একটা int variable এর pointer declare করো, pointer দিয়ে variable এর মান পরিবর্তন করো।  
- struct method pointer receiver ব্যবহার করে তার ফিল্ড update করো।  

---

### ৫. Arrays, Slices, Maps

- array এবং slice এর মধ্যে পার্থক্য কী?  
- slice এর capacity এবং length কী?  
- map কি এবং কিভাবে declare ও ব্যবহার করতে হয়?  

**Practice:**  
- একটা slice বানাও, এতে ৫টা integer যোগ করো, তারপর নতুন ভ্যালু append করো।  
- একটা map বানাও যেখানে key হবে string এবং value হবে int। তাতে কিছু ডাটা যোগ করো ও print করো।  

---

### ৬. Structs এবং Interfaces

- struct কী?  
- interface কী? Go তে interface এর zero value কী?  
- implicit interface implementation কী?  

**Practice:**  
- একটা interface বানাও যেখানে একটা method থাকবে `Speak() string`  
- দুইটা struct বানাও, যেমন Dog আর Cat, যাদের `Speak()` method থাকবে ভিন্ন ভিন্ন শব্দ return করার জন্য।  
- interface দিয়ে দুই struct এর instance কে handle করো।  

---

### ৭. Concurrency (Goroutines এবং Channels)

- goroutine কি এবং কিভাবে চালানো হয়?  
- channel কি? buffered আর unbuffered channel পার্থক্য?  
- select statement কী?  

**Practice:**  
- একটা goroutine চালাও যেটা background এ count করবে ১ থেকে ৫ পর্যন্ত।  
- একটা buffered channel বানাও এবং তাতে ডাটা পাঠাও ও রিসিভ করো।  
- select ব্যবহার করে দুই channel থেকে ডাটা নাও।  

---

### ৮. Error Handling

- Go তে error handling কিভাবে হয়?  
- `error` interface কী?  
- panic এবং recover কী?  

**Practice:**  
- একটা ফাংশন বানাও যেটা divide operation করবে এবং division by zero error handle করবে।  
- custom error type তৈরি করো।  

---

### ৯. Packages এবং Modules

- Go তে package structure কেমন?  
- module কী? go.mod ফাইলের কাজ কী?  
- কিভাবে অন্য package import এবং ব্যবহার করা হয়?  

**Practice:**  
- একটা নতুন package বানাও এবং সেখানে একটা function রাখো।  
- main package থেকে ঐ function কল করো।  

---

### ১০. Generics (Go 1.18+ Feature)

- Go তে generic functions এবং types কিভাবে তৈরি করা হয়?  
- type constraints কী?  
- কোন পরিস্থিতিতে generic ব্যবহার করা উচিত?  
