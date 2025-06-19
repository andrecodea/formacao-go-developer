<div align="center">
<h1>Go Developer Training Challenges</h1>

<p>This repository stores the challenges I completed during the **Go Developer Training at DIO**.</p>

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
</div>

## 🎯 About the Training

Training taught by professor **Tenille Martins** on the DIO platform, focused on developing practical Go skills through progressive challenges.

## 🚀 Challenges

### Challenge 1: Thermometric Converter from Kelvin to Celsius | [repo](challenge1-temperature-converter.go)

- **Objective:** Temperature scale conversion with Go
- **Challenge:** Develop a program to convert water's boiling point from Kelvin to Celsius
- **Solution:** I created a program that converts any temperature from Kelvin to Celsius and automatically identifies water's boiling and freezing points with specific messages

### Challenge 2: Numbers Divisible by 3 Detector | [repo](challenge2-divisible-by-three.go)

- **Objective:** Detect numbers divisible by 3 in Go
- **Challenge:** Implement detection using `for` loop and modulo operator (`%`)
- **Solution:** I developed a function that identifies numbers divisible by 3 within a determined range using repetition structures and mathematical operations

### Challenge 3: PinPan | [repo](challenge3-pinpan.go)

- **Objective:** Implement a custom "FizzBuzz" style program
- **Challenge:** Display "Pin" for numbers divisible by 3 and "Pan" for those divisible by 5
- **Solution:** I created a function that combines the verifications from previous challenges, detecting divisibility by 3 and 5 simultaneously using `for` loop and modulo operator (`%`)

### Challenge 4: Ping Pong | [repo](challenge4-pingpong-concurrency.go)

- **Objective:** Implement a pingpong program using concurrency
- **Challenge:** Implement concurrency to print "ping" and "pong" one after the other, using goroutines, and channels.
- **Solution:** I created two channels that send "ping" and "pong", had a receiving end in a loop that prints the message and sleeps for 1 second. Aditionally, I used sync to syncronize both goroutines and prevent the main goroutine from ending the program without any returns, using WaitGroup functions `Add()` , `Wait()` and `Done()`.

## 📚 Key Learnings

- Go language fundamentals
- Control and repetition structures
- Mathematical and logical operations
- Programming best practices
- Algorithmic problem solving

## 🤝 Contributions

Feel free to explore the code, suggest improvements, and point out mistakes! After all, I'm just learning.


## 📄 License
This project is licensed under the MIT License - see the [LICENSE](LICENSE.txt) file for details.

---
<div align="center">
  <b>⭐ Liked the project? Leave a star on the repository!</b>
</div>
