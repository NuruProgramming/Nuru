# Pakeji Hesabu (Math Package)

Pakeji Hesabu is a math package written in pure Nuru by [VictorKariuki](https://github.com/VictorKariuki).

This package provides various mathematical functions and constants implemented in nuru programming language. It includes methods for `trigonometric functions`, `logarithmic functions`, `array operations`, and `utility functions`.


## Usage
To use the `pakeji hesabu` package follow the steps below:

1. Copy the `hesabu.nr` file and any required third-party package files into the same directory as your project.

2. Ensure that the package file names end with the `.nr` extension and match the package names. For example, if the package name is `hesabu`, the corresponding file name should be `hesabu.nr`.

3. You can directly import the `hesabu.nr` package and any required third-party packages in your Nuru code using the `tumia` keyword. For example:

   ```nuru
   tumia "hesabu"
   ```
   Example of calling the package methods:
   ```nuru
   andika(hesabu.e())
## What is in
This package covers a wide range of mathematical operations, including `basic arithmetic`, `trigonometry`, `exponential and logarithmic functions`, `rounding and comparison operations`, as well as some `utility and array operations`.

The methods provided in the `hesabu` package can be classified into different categories based on their functionalities. Here is a classification of the methods:

1. Trigonometric Functions:
   - `cos(x)`
   - `sin(x)`
   - `tan(x)`
   - `acos(x)`
   - `asin(x)`
   - `atan(x)`

2. Hyperbolic Functions:
   - `cosh(x)`
   - `sinh(x)`
   - `tanh(x)`
   - `acosh(x)`
   - `asinh(x)`
   - `atanh(x)`

3. Exponential and Logarithmic Functions:
   - `exp(x)`
   - `expm1(x)`
   - `log(x)`
   - `log10(x)`
   - `log1p(x)`

4. Other Mathematical Functions:
   - `abs(namba)`
   - `ceil(x)`
   - `floor(x)`
   - `sqrt(x)`
   - `cbrt(x)`
   - `root(x, n)`
   - `hypot(values)`
   - `factorial(n)`

5. Rounding and Comparison Functions:
   - `round(x, method)`
   - `max(numbers)`
   - `min(numbers)`

6. Utility Functions:
   - `sign(x)`
   - `isNegative(num)`
   - `isInteger(num)`
   - `getIntegerPart(num)`

7. Array and List Operations:
   - `list(first, last, interval)`
   - `reduce(iterator, callback, initialValue)`


### 1. Constants:
   - **PI**: Represents the mathematical constant `π`.
   - **e**: Represents `Euler's Number`.
   - **phi**: Represents the `Golden Ratio`.
   - **ln10**: Represents the `natural logarithm of 10`.
   - **ln2**: Represents the `natural logarithm of 2`.
   - **log10e**: Represents the `base 10 logarithms` of Euler's number `(e)`.
   - **log2e**: Represents the `base 2 logarithm` of Euler's number` (e)`.
   - **sqrt1_2**: Represents the `square root` of `1/2`.
   - **sqrt2**: Represents the `square root` of `2`.
   - **sqrt3**: Represents the `square root` of `3`.
   - **sqrt5**: Represents the `square root` of `5`.
   - **EPSILON**: Represents a small value (2.220446049250313e-16).

### 2. Methods:

1. **abs(namba)**
   - Description: Calculates the absolute value of a number.
   - Example: `hesabu.abs(-42)` returns `42`.

2. **acos(x)**
   - Description: Calculates the arccosine of a number.
   - Example: `hesabu.acos(0.5)` returns `1.0471975511965979`.

3. **acosh(x)**
   - Description: Calculates the inverse hyperbolic cosine of a number.
   - Example: `hesabu.acosh(2)` returns `1.3169578969248166`.

4. **asin(x)**
   - Description: Calculates the arcsine of a number using the Taylor series.
   - Example: `hesabu.arcsin(0.5)` returns `0.5235987755982989`.

5. **asinh(x)**
   - Description: Calculates the inverse hyperbolic sine of a number.
   - Example: `hesabu.arsinh(2)` returns `1.4436354751788103`.

6. **atan(x)**
   - Description: Calculates the arctangent of a number using the Taylor series.
   - Example: `hesabu.atan(1)` returns `0.7853981633974485`.

7. **atan2(y, x)**
   - Description: Calculates the arctangent of the quotient of its arguments.
   - Example: `hesabu.atan2(1, 1)` returns `0.7853981633974483`.

8. **atanh(x)**
   - Description: Calculates the inverse hyperbolic tangent of a number.
   - Example: `hesabu.atanh(0.5)` returns `0.5493061443340549`.

9. **cbrt(x)**
   - Description: Calculates the cube root of a number.
   - Example: `hesabu.cbrt(8)` returns `2`.

10. **root(x, n)**
    - Description: Calculates the nth root of a number using the Newton-Raphson method.
    - Example: `hesabu.root(27, 3)` returns `3`.

11. **ceil(x)**
    - Description: Rounds up to the smallest integer greater than or equal to a given number.
    - Example: `hesabu.ceil(4.3)` returns `5`.

12. **cos(x)**
    - Description: Calculates the cosine of an angle in radians using the Taylor series.
    - Example: `hesabu.cos(5)` returns `0.28366218546322464`.

13. **cosh(x)**
    - Description: Calculates the hyperbolic cosine of a number.
    - Example: `hesabu.cosh(5)` returns `74.20994842490012`.

14. **exp(x)**
    - Description: Calculates the value of Euler's number raised to the power of a given number.
    - Example: `hesabu.exp(2)` returns `7.389056098930649`.

15. **expm1(x)**
    - Description: Calculates Euler's number raised to the power of a number minus 1.
    - Example: `hesabu.expm1(1)` returns `1.7182818284590455`.

16. **floor(x)**
    - Description: Rounds down to the largest integer less than or equal to a given number.
    - Example: `hesabu.floor(4.7)` returns `4`.

17. **hypot(values)**
    - Description: Calculates the square root of the sum of squares of the given values.
    - Example: `hesabu.hypot([3, 4])` returns `5`.

18. **log(x)**
    - Description: Calculates the natural logarithm of a number.
    - Example: `hesabu.log(2)` returns `0.69314718056`.

19. **log10(x)**
    - Description: Calculates the base 10 logarithm of a number.
    - Example: `hesabu.log10(100)` returns `1.9999999999573126`.

20. **log1p(x)**
    - Description: Calculates the natural logarithm of 1 plus the given number.
    - Example: `hesabu.log1p(1)` returns `0.6931471805599451`.

21. **log2(x)**
    - Description: Calculates the base 2 logarithm of a number.
    - Example: `hesabu.log2(8)` returns `3`.

22. **max(numbers)**
    - Description: Finds the maximum value in a list of numbers.
    - Example: `hesabu.max([4, 2, 9, 5])` returns `9`.

23. **min(numbers)**
    - Description: Finds the minimum value in a list of numbers.
    - Example: `hesabu.min([4, 2, 9, 5])` returns `2`.

24. **round(x, method)**
    - Description: Rounds a number to the nearest integer using the specified method.
    - supported methods:
        - "rpi" (round to the nearest integer using the principle of rounding half to the nearest even)
        - "rni" (round to the nearest integer using the principle of rounding half away from zero)
        - "ri" (round to the nearest integer using the standard rounding method)
        - An invalid method results in returning NaN (Not a Number)
    - Example: `hesabu.round(4.6, "rpi")` returns `5`.

25. **sign(x)**
    - Description: Determines the sign of a number.
    - Example: `hesabu.sign(-5)` returns `-1`.

26. **sin(x)**
    - Description: Calculates the sine of an angle in radians using the Taylor series.
    - Example: `hesabu.sin(1)` returns `0.8414709848078965`.

27. **sinh(x)**
    - Description: Calculates the hyperbolic sine of a number.
    - Example: `hesabu.sinh(0)` returns `0`.

28. **sqrt(x)**
    - Description: Calculates the square root of a number.
    - Example: `hesabu.sqrt(4)` returns `2`.

29. **tan(x)**
    - Description: Calculates the tangent of an angle in radians.
    - Example: `hesabu.tan(1)` returns `1.557407724654902`.

30. **tanh(x)**
    - Description: Calculates the hyperbolic tangent of a number.
    - Example: `hesabu.tanh(0)` returns `0`.

31. **factorial(n)**
    - Description: Calculates the factorial of a number.
    - Example: `hesabu.factorial(5)` returns `120`.

32. **isNegative(num)**
    - Description: Checks if a number is negative.
    - Example: `hesabu.isNegative(-5)` returns `kweli`.

33. **isInteger(num)**
    - Description: Checks if a number is an integer.
    - Example: `hesabu.isInteger(4.5)` returns `sikweli`.

34. **getIntegerPart(num)**
    - Description: Gets the integer part of a number.
    - Example: `hesabu.getIntegerPart(4.5)` returns `4`.

35. **list(first, last, interval)**
    - Description: Creates a list of numbers with the specified interval between them.
    - Example: `hesabu.list(1, 5, 1)` returns `[1, 2, 3, 4]`.

36. **reduce(iterator, callback, initialValue)**
      - Description: Reduces the elements of an array to a single value using a specified callback function.
      - Example: `hesabu.reduce([1, 2, 3, 4], [callback function], 0)`
      ```s
        fanya callback = unda(accumulator, currentValue){
            rudisha accumulator + currentValue;
        }

        andika(hesabu.reduce([1, 2, 3, 4], callback, 0)) \\ returns 10.
### Contributing

Contributions to the `pakeji hesabu` package are welcome. If you have any improvements or bug fixes, feel free to create a pull request.

### License

This package is available under the MIT License. See the [LICENSE](LICENSE) file for more information.