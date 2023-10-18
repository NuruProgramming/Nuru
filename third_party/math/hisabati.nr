pakeji hisabati{
    //CONSTRUCTOR METHOD    
    andaa = unda() {}

    // Constants
    // π (Pi)
    PI = unda() {
        rudisha 3.141592653589793;
    }

    // e (Euler's Number)
    e = unda() {
        rudisha 2.718281828459045;
    }

    // φ (Phi, Golden Ratio)
    phi = unda() {
        rudisha 1.618033988749895;
    }

    // Natural logarithm of 10
    ln10 = unda() {
        rudisha 2.302585092994046;
    }

    // Natural logarithm of 2
    ln2 = unda() {
        rudisha 0.6931471805599453;
    }

    // Base 10 logarithm of Euler's number (e)
    log10e = unda() {
        rudisha 0.4342944819032518;
    }

    // Base 2 logarithm of Euler's number (e)
    log2e = unda() {
        rudisha 1.4426950408889634;
    }

    // √1/2 (equivalent to 1 / sqrt(2))
    sqrt1_2 = unda() {
        rudisha 0.7071067811865476;
    }

    // √2 (Square Root of 2)
    sqrt2 = unda() {
        rudisha 1.414213562373095;
    }

    // √3 (Square Root of 3)
    sqrt3 = unda() {
        rudisha 1.732050807568877;
    }

    // √5 (Square Root of 5)
    sqrt5 = unda() {
        rudisha 2.236067977499790;
    }

    // @.EPSILON
    EPSILON = unda() {
        rudisha 2.220446049250313e-16;
    }

    // Methods
    //abs(namba), calculates the absolute value of a number.
    abs = unda(namba){
        kama(namba < 0){
            rudisha - 1 * namba;
        }

        rudisha namba;
    }

    //acos(x), calculates the arccosine of a number.
    acos = unda(x) {
        kama(x < -1 || x > 1) {
            rudisha 0;
        }

        // Define the precision for the approximation.
        fanya precision = 1e-10;

        // Initial guess for the angle in radians (between 0 and π).
        fanya angle = @.PI / 4;

        wakati(true) {
            fanya cosAngle = @.cos(angle);
            fanya error = @.abs(cosAngle - x);

            kama(error < precision) {
                rudisha angle;
            }

            // Update the angle using the Newton-Raphson method.
            angle -= (cosAngle - x) / (-@.sin(angle));
        }
    }

    //acosh(x), calculates the inverse hyperbolic cosine of a number.
    acosh = unda(x) {
        kama(x < 1) {
            rudisha 0;
        }

        rudisha @.log(x + @.sqrt(x * x - 1));
    }

    //arcsin(x), calculates the arcsine of a number using the Taylor series.
    arcsin = unda(x) {
        kama(x < -1 || x > 1) {
            rudisha 0;
        }

        fanya maxIterations = 50; // Maximum number of iterations
        fanya result = 0;

        fanya n = 0;
        wakati(n < maxIterations) {
            fanya numerator = (-1) ** n * x ** (2 * n + 1);
            fanya denominator = @.factorial(2 * n + 1); // You'll need to implement a factorial function

            result += numerator / denominator;
            n++;
        }

        rudisha result;
    }

    //arsinh(x), calculates the inverse hyperbolic sine of a number.
    arsinh = unda(x) {
        // Calculate arsinh using the formula: arsinh(x) = ln(x + sqrt(x^2 + 1))
        kama(x >= 0) {
            rudisha @.log(x + @.sqrt(x * x + 1));
        } sivyo {
            // For negative values, arsinh(x) = -arsinh(-x)
            rudisha - @.log(-x + @.sqrt(x * x + 1));
        }
    }

    //atan(x), calculates the arctangent of a number using the Taylor series.
    atan = unda(x) {
        kama(x == 0) {
            rudisha 0;
        } // arctan(0) is 0
        kama(x == Infinity) {
            rudisha @.PI / 2;
        }
        kama(x == -Infinity) {
            rudisha - @.PI / 2;
        } // arctan(-Infinity) is -π/2

        // Use the Taylor series expansion for arctan(x)
        // arctan(x) = x - (x^3) / 3 + (x^5) / 5 - (x^7) / 7 + ...
        fanya n = 3;
        fanya result = 0;
        fanya term = x * x;
        fanya sign = -1;

        wakati(true) {
            fanya currentTerm = sign * (term / n);

            kama(currentTerm == 0) {
                vunja
            }

            result += currentTerm;
            n += 2;
            sign = -sign;
            term *= x * x;
        }

        rudisha result;
    }

    //atanh(x), calculates the inverse hyperbolic tangent of a number.
    atan2 = unda(y, x) {
        kama(x > 0) {
            rudisha @.atan(y / x);
        } au kama(x < 0 && y >= 0) {
            rudisha @.atan(y / x) + @.PI;
        } au kama(x < 0 && y < 0) {
            rudisha @.atan(y / x) - @.PI;
        } au kama(x == 0 && y > 0) {
            rudisha @.PI / 2;
        } au kama(x == 0 && y < 0) {
            rudisha - @.PI / 2;
        } au kama(x == 0 && y == 0) {
            rudisha NaN; // Undefined
        }
    }

    //atanh(x), calculates the inverse hyperbolic tangent of a number.
    atanh = unda(x) {
        kama(x < -1 || x > 1) {
            rudisha 0;
        }
        rudisha 0.5 * @.log((1 + x) / (1 - x));
    }

    //cbrt(x), calculates the cube root of a number.
    cbrt = unda(x) {
        kama(x >= 0) {
            rudisha @.root(x, 3);
        } sivyo {
            rudisha - @.root(-x, 3);
        }
    }

    //root(x, n), calculates the nth root of a number using the Newton-Raphson method.
    root = unda(x, n) {
        fanya guess = x / 2; // Initial guess
        fanya tolerance = 1e-10; // Tolerance for convergence

        wakati(true) {
            fanya nextGuess = ((n - 1) * guess + x / guess ** n - 1) / n;
            kama(@.abs(nextGuess - guess) < tolerance) {
                rudisha nextGuess;
            }
            guess = nextGuess;
        }
    }

    //ceil(x), rounds up to the smallest integer greater than or equal to a given number.
    ceil = unda(x) {
        kama(x >= 0) {
            kama(x % 1 == 0) {
            rudisha x; // x is already an integer
            }
            rudisha floor(x) + 1;
        } sivyo {
            rudisha - floor(abs(x));
        }
    }

    //cos(x), calculates the cosine of an angle in radians using the Taylor series.
    cos = unda(x, terms = 10) {
        // Initialize the result
        fanya n = 0;
        fanya result = 0;

        wakati(n < terms) {
            // Calculate the numerator and denominator for the nth term
            fanya numerator = 0;
            kama(n % 2 == 0) {
                numerator = x ** (2 * n + 1);
            } sivyo {
                numerator = -x ** (2 * n + 1);
            }
            fanya denominator = @.factorial(2 * n);

            // Add the nth term to the result
            result += numerator / denominator;

            n++;
        }

        rudisha result;
    }

    //cosh(x), calculates the hyperbolic cosine of a number.
    cosh = unda(x) {
        fanya eToX = @.exp(x);
        fanya eToMinusX = @.exp(-x);
        rudisha(eToX + eToMinusX) / 2;
    }

    //exp(x), calculates the value of Euler's number raised to the power of a given number.
    exp = unda(x, precision = 15) {
        fanya result = 1;
        fanya term = 1;
        fanya i = 1;

        wakati(i <= precision) {
            term *= x / i;
            result += term;
            i++;
        }


        rudisha result;
    }

    //expm1(x), calculates the value of Euler's number raised to the power of a given number minus 1.
    expm1 = unda(x) {
        fanya epsilon = 1e-15; // A small value to improve accuracy
        fanya result = 0;
        fanya term = x;
        fanya n = 1;

        wakati(@.abs(term) > epsilon) {
            result += term;
            n++;
            term *= x / n;
        }

        rudisha result;
    }

    //floor(x), rounds down to the largest integer less than or equal to a given number.
    floor = unda(x) {
        kama(x >= 0) {
            rudisha x - (x % 1);
        } sivyo {
            rudisha x - (1 + x % 1);
        }
    }

    //hypot(values), calculates the square root of the sum of squares of the given values.
    hypot = unda(values) {
        // Calculate the sum of squares of the values
        fanya exp = unda(acc, value){
            rudisha acc + value ** 2;
        }

        fanya sumOfSquares = @.reduce(values, exp, 0);

        // Calculate the square root of the sum of squares
        fanya result = @.sqrt(sumOfSquares);

        rudisha result;
    }

    //log(x), calculates the natural logarithm of a number.
    log = unda(x) {
        kama(x <= 0) {
            rudisha 0;
        }

        fanya approx = 0;
        fanya n = 50;
        fanya i = 1;

        wakati(i <= n) {
            approx += (1 / i) * ((x - 1) / (x + 1)) ** (2 * i - 1);
            i++;
        }


        rudisha 2 * approx;
    }

    //log10(x), calculates the base 10 logarithm of a number.
    log10 = unda(x) {
        kama(x <= 0) {
            rudisha 0;
        }

        // Calculate natural logarithm and divide by the natural logarithm of 10
        rudisha this.log(x) / this.log(10);
    }

    //log1p(x), calculates the natural logarithm of 1 plus the given number.
    log1p = unda(x) {
        kama(x < -1) {
            rudisha 0;
        }

        kama(x == -1) {
            rudisha(-Infinity);
        }

        kama(x == Infinity) {
            rudisha Infinity;
        }

        kama(x == 0) {
            rudisha 0;
        }

        // Use the formula: ln(1 + x) = x - (x^2)/2 + (x^3)/3 - (x^4)/4 + ...
        fanya result = 0;
        fanya term = x;
        fanya i = 2;

        wakati(@.abs(term) > @.EPSILON) {
            result += term;
            term *= -x / i;
            i++;
        }

        rudisha result;
    }

    //log2(x), calculates the base 2 logarithm of a number.
    log2 = unda(x) {
        kama(x <= 0) {
            rudisha 0;
        }

        fanya result = 0;
        fanya currentValue = x;

        wakati(currentValue > 1) {
            currentValue /= 2;
            result++;
        }

        rudisha result;
    }

    //max(numbers), finds the maximum value in a list of numbers.
    max = unda(numbers) {
        // Initialize a variable to store the largest number
        fanya largest = -Infinity;

        // Iterate through the numbers and update 'largest' kama a larger number is found
        kwa num ktk numbers {
            kama(num > largest) {
                largest = num;
            }
        }

        // rudisha the largest number (or -Infinity kama there are no parameters)
        rudisha largest;
    }

    //min(numbers), finds the minimum value in a list of numbers.
    min = unda(numbers) {
        kama(numbers.length == 0) {
            rudisha Infinity;
        }

        fanya minVal = numbers[0];

        fanya i = 1;
        wakati(i < numbers.length) {
            kama(numbers[i] < minVal) {
                minVal = numbers[i];
            }
            i++;
        }

        rudisha minVal;
    }

    //round(x, method), rounds a number to the nearest integer using the specified method.
    round = unda(x, method = "ri") {
        kama(method == "rpi") {
            rudisha floor(x + 0.5);
        } au kama(method == "rni") {
            rudisha ceiling(x - 0.5);
        } au kama(method == "ri") {
            kama(x >= 0){
                rudisha floor(x + 0.5);
            }sivyo{
                rudisha ceiling(x - 0.5);
            }
        } sivyo {
            rudisha NaN; // Invalid method
        }
    }

    //sign(x), determines the sign of a number.
    sign = unda(x) {
        kama(x == 0 || x == -0) {
            rudisha x;
        } au kama(x > 0) {
            rudisha 1;
        } sivyo {
            rudisha - 1;
        }
    }

    //sin(x), calculates the sine of an angle in radians using the Taylor series.
    sin = unda(x, terms = 10) {
        // Initialize the result
        fanya n = 0;
        fanya result = 0;

        wakati(n < terms) {
            // Calculate the numerator and denominator for the nth term
            fanya numerator = 0;
            kama(n % 2 == 0) {
                numerator = x ** (2 * n + 1);
            } sivyo {
                numerator = -x ** (2 * n + 1);
            }

            fanya denominator = @.factorial(2 * n + 1);

            // Add the nth term to the result
            result += numerator / denominator;

            n++;
        }

        rudisha result;
    }

    //sinh(x), calculates the hyperbolic sine of a number.
    sinh = unda(x) {
        // sinh(x) = (e^x - e^(-x)) / 2
        fanya eToX = @.exp(x);
        fanya eToMinusX = @.exp(-x);
        rudisha(eToX - eToMinusX) / 2;
    }

    //sqrt(x), calculates the square root of a number.
    sqrt = unda(x) {
        kama(x < 0) {
            rudisha 0;
        }

        // Initial guess for the square root (you can choose a better initial guess)
        fanya guess = x / 2;
        fanya tolerance = 1e-7; // Tolerance for approximation

        wakati(true) {
            fanya nextGuess = 0.5 * (guess + x / guess);

            // Check kama the guess is close enough to the actual square root
            kama(@.abs(nextGuess - guess) < tolerance) {
                rudisha nextGuess;
            }

            guess = nextGuess;
        }
    }

    //tan(x), calculates the tangent of an angle in radians.
    tan = unda(x) {
        fanya sineX = @.sine(x);
        fanya cosineX = @.sqrt(1 - sineX * sineX);

        kama(cosineX == 0) {
            rudisha 0;
        }

        rudisha sineX / cosineX;
    }

    //tanh(x), calculates the hyperbolic tangent of a number.
    tanh = unda(x) {
        kama(x == Infinity) {
            rudisha 1;
        } au kama(x == -Infinity) {
            rudisha - 1;
        } sivyo {
            fanya expX = @.exp(x);
            fanya expNegX = @.exp(-x);
            rudisha(expX - expNegX) / (expX + expNegX);
        }
    }

    // utility methods
    //factorial(n), calculates the factorial of a number.
    factorial = unda(n) {
        kama(n == 0){ 
        rudisha 1;
        };
        fanya result = 1;
        fanya i = 1;

        wakati(i <= n) {
            result *= i;
            i++;
        }

        rudisha result;

    }

    //isNegative(num), checks if a number is negative.
    isNegative = unda(num) {
        rudisha num < 0;
    }

    //isInteger(num), checks if a number is an integer.
    isInteger = unda(num) {
        rudisha num == @.floor(num);
    }

    //getIntegerPart(num), gets the integer part of a number.
    getIntegerPart = unda(num) {
        // Handle negative numbers separately
        kama(@.isNegative(num)) {
            // For negative numbers, we subtract the absolute value of the fractional part from 1
            rudisha - (@.ceil(-num) - 1);
        } sivyo {
            // For positive numbers, we simply truncate the fractional part
            rudisha @.floor(num);
        }
    }

    //Arrray Methods
    //list(first, last, interval), creates a list of numbers with the specified interval between theM.
    list = unda(first, last, interval){
        fanya list = [first];
        fanya i = first + interval;
        wakati(i < last){
            list.sukuma(i);
            i += interval;
        }
        rudisha list;
    }

    //reduce(iterator, callback, initialValue), reduces the elements of an array to a single value using a specified callback function.
    reduce = unda(iterator, callback, initialValue) {
        fanya accumulator = initialValue;

        kwa thamani ktk iterator {
            accumulator = callback(accumulator, thamani);
        }

        rudisha accumulator;
    }
}
