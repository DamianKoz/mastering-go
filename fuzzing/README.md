# Fuzzing

Fuzzy testing is a great way to increase the resilience and robustness of your code.

Fuzzy tests generate inputs for your programs, which might lead to panics, bugs or data races. Thereby revelealing possible problems and security issues in your code.

It's worth noting that effective fuzzing necessitates an iterative approach, continuously improving the seed corpus and refining the fuzz test function to uncover new issues over time.

Fuzzing uses semi-random data mutations to handle more edge-cases that a developer might not have thought of in the normal test cases.

A Fuzz tests contains (most importantly) of a fuzz test, a seed corpus and the test target.

- The **Fuzz test** is a function with the key "Fuzz" in the beginning like **FuzzReverse**.
- The **seed corpus** is a set of inputs that are used to test your code. You can add something to the corpus via `f.Add("cat", big.NewInt(1341))`. The corpus is also used by the fuzzing engine to generate more inputs similar to what the programmer provided.
- The **test target** is, as the name already suggests, the target to be tested. It is a method call to (*testing.F).Fuzz which accepts a *testing.T as the first parameter, followed by the fuzzing arguments.

You can find more information about fuzzing at these resources:

- https://go.dev/security/fuzz/
- https://go.googlesource.com/proposal/+/master/design/draft-fuzzing.md#seed-corpus
- https://go.dev/security/best-practices

Fuzzing can significantly contribute to the overall reliability and security of your software by identifying and mitigating potential weaknesses that traditional testing may overlook. By incorporating fuzzing into your testing strategy, you can enhance the quality of your codebase and bolster your application's resilience against unforeseen scenarios.
