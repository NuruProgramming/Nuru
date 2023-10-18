# Math Package



## Explanation
This math package provides a collection of mathematical functions and constants implemented in the Nuru programming language. These functions cover a wide range of mathematical operations, including trigonometric functions, logarithmic functions, and other common mathematical operations. Below is a detailed list of all methods along with their descriptions and examples:

1. **abs(namba)**
   - Description: Calculates the absolute value of a number.
   - Example: `Hisabati.abs(-42)` returns `42`.

2. **acos(x)**
   - Description: Calculates the arccosine (inverse cosine) of a number in radians.
   - Example: `Hisabati.acos(0.5)` returns the arccosine of 0.5.

3. **acosh(x)**
   - Description: Calculates the inverse hyperbolic cosine of a number.
   - Example: `Hisabati.acosh(2)` returns the inverse hyperbolic cosine of 2.

4. **arcsin(x)**
   - Description: Calculates the arcsine (inverse sine) of a number using a Taylor series.
   - Example: `Hisabati.arcsin(0.5)` returns the arcsine of 0.5.

5. **arsinh(x)**
   - Description: Calculates the inverse hyperbolic sine of a number.
   - Example: `Hisabati.arsinh(1)` returns the inverse hyperbolic sine of 1.

6. **atan(x)**
   - Description: Calculates the arctangent (inverse tangent) of a number in radians using a Taylor series.
   - Example: `Hisabati.atan(1)` returns the arctangent of 1.

7. **atan2(y, x)**
   - Description: Calculates the angle in radians between the positive x-axis and the point (x, y).
   - Example: `Hisabati.atan2(1, 1)` returns the angle for the point (1, 1).

8. **atanh(x)**
   - Description: Calculates the inverse hyperbolic tangent of a number.
   - Example: `Hisabati.atanh(0.5)` returns the inverse hyperbolic tangent of 0.5.

9. **cbrt(x)**
   - Description: Calculates the cube root of a number.
   - Example: `Hisabati.cbrt(8)` returns the cube root of 8.

10. **ceil(x)**
    - Description: Rounds up to the smallest integer greater than or equal to a given number.
    - Example: `Hisabati.ceil(4.2)` returns `5`.

11. **cos(x, terms)**
    - Description: Calculates the cosine of a number in radians using a Taylor series.
    - Example: `Hisabati.cos(1)` returns the cosine of 1.

12. **cosh(x)**
    - Description: Calculates the hyperbolic cosine of a number.
    - Example: `Hisabati.cosh(2)` returns the hyperbolic cosine of 2.

13. **exp(x, precision)**
    - Description: Calculates the value of the mathematical constant e raised to the power of x using a Taylor series.
    - Example: `Hisabati.exp(2)` returns `e^2`.

14. **expm1(x)**
    - Description: Calculates the value of e^x - 1 using a Taylor series.
    - Example: `Hisabati.expm1(1)` returns `e - 1`.

15. **floor(x)**
    - Description: Rounds down to the largest integer less than or equal to a given number.
    - Example: `Hisabati.floor(4.9)` returns `4`.

16. **hypot(values)**
    - Description: Calculates the Euclidean norm (square root of the sum of squares) of a list of values.
    - Example: `Hisabati.hypot([3, 4])` returns `5`, which is the hypotenuse of a right triangle.

17. **log(x)**
    - Description: Calculates the natural logarithm of a number using a Taylor series.
    - Example: `Hisabati.log(2)` returns the natural logarithm of 2.

18. **log10(x)**
    - Description: Calculates the base 10 logarithm of a number using the natural logarithm.
    - Example: `Hisabati.log10(100)` returns `2`.

19. **log1p(x)**
    - Description: Calculates the natural logarithm of 1 + x using a Taylor series.
    - Example: `Hisabati.log1p(0.5)` returns the natural logarithm of 1.5.

20. **log2(x)**


    - Description: Calculates the base 2 logarithm of a number.
    - Example: `Hisabati.log2(8)` returns `3`.

21. **max(numbers)**
    - Description: Returns the largest number from a list of numbers.
    - Example: `Hisabati.max([3, 7, 2, 9])` returns `9`.

22. **min(numbers)**
    - Description: Returns the smallest number from a list of numbers.
    - Example: `Hisabati.min([3, 7, 2, 9])` returns `2`.

23. **round(x, method)**
    - Description: Rounds a number to the nearest integer using different rounding methods.
    - Example: `Hisabati.round(3.6, "rpi")` rounds to the nearest integer (`4`) using "round half up."

24. **sign(x)**
    - Description: Returns the sign of a number: `1` for positive, `-1` for negative, and `0` for zero.
    - Example: `Hisabati.sign(-7)` returns `-1`.

25. **sin(x, terms)**
    - Description: Calculates the sine of a number in radians using a Taylor series.
    - Example: `Hisabati.sin(0.5)` returns the sine of 0.5.

26. **sinh(x)**
    - Description: Calculates the hyperbolic sine of a number.
    - Example: `Hisabati.sinh(1)` returns the hyperbolic sine of 1.

27. **sqrt(x)**
    - Description: Calculates the square root of a number using the Newton-Raphson method.
    - Example: `Hisabati.sqrt(16)` returns `4`.

28. **tangent(x)**
    - Description: Calculates the tangent of a number in radians.
    - Example: `Hisabati.tangent(1)` returns the tangent of 1.

29. **tanh(x)**
    - Description: Calculates the hyperbolic tangent of a number.
    - Example: `Hisabati.tanh(0.5)` returns the hyperbolic tangent of 0.5.

30. **isNegativeZero(num)**
    - Description: Checks if a number is negative zero (0 with a negative sign).
    - Example: `Hisabati.isNegativeZero(-0)` returns `true`.

31. **factorial(n)**
    - Description: Calculates the factorial of a number.
    - Example: `Hisabati.factorial(5)` returns `120`.

32. **pow(base, exponent)**
    - Description: Calculates the power of a number.
    - Example: `Hisabati.pow(2, 3)` returns `8`.

33. **isNegative(num)**
    - Description: Checks if a number is negative.
    - Example: `Hisabati.isNegative(-5)` returns `true`.

34. **isInteger(num)**
    - Description: Checks if a number is an integer.
    - Example: `Hisabati.isInteger(42)` returns `true`.

35. **getIntegerPart(num)**
    - Description: Gets the integer part of a number.
    - Example: `Hisabati.getIntegerPart(5.8)` returns `5`.

36. **list(first, last, interval)**
    - Description: Creates a list of numbers within a specified range with a given interval.
    - Example: `Hisabati.list(0, 10, 2)` returns `[0, 2, 4, 6, 8]`.

37. **square(n, i, j)**
    - Description: Finds the square root of a number using a method that iteratively narrows down the root.
    - Example: `Hisabati.square(16)` returns `4`.

Feel free to use this package for your mathematical calculations and applications.