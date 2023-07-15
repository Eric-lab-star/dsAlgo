---
date created: 2023-06-04 08:19
---

Asymptotic notation is a way to describe the running time or space complexity of an algorithm based on the input size.

1. Big O notations
   1. Provides upper bounds
   2. Represents worst case scenario
2. Omega notation
   1. Lower bounds
   2. best case scenario
3. Theta notation
   1. Both upper and lower bound
   2. Average-case scenario

Asymptotic Notation Features

1. Outcome does not match to real running time or space usage
2. Only useful for comparing among several algorithm

## Why performance anlaysis?

Performance is speed. Speed can be seen as money to but user-friendliness, modularity, security, maintainability.

## How to study efficient of algorithms?

Run various test inputs while recording the time spent for each execution.

## Given two algorithms for a task, how do we find out which one is better?

Naive approach is to run two programs for various inputs and find which one takes less time.

Problems of naive approach

- different input may cause differnt result
- machine may cause effect on result

Asymptotic Analysis is alternative for navie approach.

## Does Asymptotic Analysis always work?

No, 100nLogn and 2nLogn time are aysmptotically same. Moreover, asymptotically slower algorithm can be suitable in some situations.

## Advantages And Disadvantages

Advantage

1. high-level understaning of how an algorithm performs with respect to input size can be aquired.
2. Useful tool for comparing the efficiency of differnet algorithms and selecting the best one for a specific problem.
3. It helps in predicting how an algorithm will perform on larger input sizes, which is essential ofr real-world applications.
4. Aymptotic analysis is relatively easy to perform and requries only basic mathmetical skills

Disadvantages:

1. Asymptotic analysis does not provide accurate running time or space usage of an algorithm
2. It assumes that the input size is the only factor that affects an algorithm's performance, which is not always the case in practice.
3. Asymptotic analysis can sometimes be misleading, as two algorithms with the same asymptotic complexity may have different actual running times or space usage.
4. It is not always straightforward to determine the best asymptotic complexity for an algorithm, as there may be trade-offs between time and space complexity.
