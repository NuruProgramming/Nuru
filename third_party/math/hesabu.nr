tumia hisabati

pakeji hesabu{
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
        rudisha 0.0000000000000002220446049250313;
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
        kama (x < -1 || x > 1) {
            rudisha "NaN";
        }

        fanya EPSILON = 1*10.0**-10; // Small value for precision

        fanya acosRecursive = unda(guess) {
            fanya f = cos(guess) - x;
            fanya fPrime = -sin(guess);
            fanya nextGuess = guess - f / fPrime;

            kama (abs(nextGuess - guess) < EPSILON) {
                rudisha nextGuess;
            }

            rudisha acosRecursive(nextGuess);
        }

        rudisha acosRecursive(hisabati.PI() / 2); // Initial guess for acos
    }

    //acosh(x), calculates the inverse hyperbolic cosine of a number.
    acosh = unda(x) {
        kama(x < 1) {
            rudisha 0;
        }

        rudisha log(x + sqrt(x * x - 1));
    }

    //asin(x), calculates the arcsine of a number using the Newton Method.
    asin = unda(x) {
        kama (x < -1 || x > 1) {
            rudisha "NaN";
        }

        fanya maxIterations = 50; // Maximum number of iterations

        fanya newtonAsin = unda(guess, prev, iterations) {
            fanya next = guess - (sin(guess) - x) / cos(guess);

            kama (abs(next - prev) < hisabati.EPSILON() || iterations >= maxIterations) {
                rudisha next;
            }

            rudisha newtonAsin(next, guess, iterations + 1);
        }

        rudisha newtonAsin(x, 1, 0);
    }


    //asinh(x), calculates the inverse hyperbolic sine of a number.
    asinh = unda(x) {
        // Calculate arsinh using the formula: arsinh(x) = ln(x + sqrt(x^2 + 1))
        kama(x >= 0) {
            rudisha log(x + sqrt(x * x + 1));
        } sivyo {
            // For negative values, arsinh(x) = -arsinh(-x)
            rudisha - log(-x + sqrt(x * x + 1));
        }
    }

    //atan(x), calculates the arctangent of a number using the Taylor series.
    atan = unda(x) {
        fanya EPSILON = 1*10.0**-10; // Small value for precision

        fanya atanRecursive = unda(guess) {
            fanya f = tan(guess) - x;
            fanya fPrime = 1 / (cos(guess) * cos(guess));
            fanya nextGuess = guess - f / fPrime;

            kama (abs(nextGuess - guess) < EPSILON) {
                rudisha nextGuess;
            }

            rudisha atanRecursive(nextGuess);
        }

        rudisha atanRecursive(x); // Initial guess for atan
    }

    //atanh(x), calculates the inverse hyperbolic tangent of a number.
    atan2 = unda(y, x) {
        kama(x > 0) {
            rudisha atan(y / x);
        } au kama(x < 0 && y >= 0) {
            rudisha atan(y / x) + hisabati.PI();
        } au kama(x < 0 && y < 0) {
            rudisha atan(y / x) - hisabati.PI();
        } au kama(x == 0 && y > 0) {
            rudisha hisabati.PI() / 2;
        } au kama(x == 0 && y < 0) {
            rudisha - hisabati.PI() / 2;
        } au kama(x == 0 && y == 0) {
            rudisha "NaN"; // Undefined
        }
    }

    //atanh(x), calculates the inverse hyperbolic tangent of a number.
    atanh = unda(x) {
        kama(x < -1 || x > 1) {
            rudisha 0;
        }
        rudisha 0.5 * log((1.0 + x) / (1.0 - x));
    }

    //cbrt(x), calculates the cube root of a number.
    cbrt = unda(x) {
        kama(x == 0) {
            rudisha 0;
        }
        
        kama(x >= 0) {
            rudisha root(x, 3);
        } sivyo {
            rudisha - root(-x, 3);
        }
    }

    //root(x, n), calculates the nth root of a number using the Newton-Raphson method.
    root =  unda(x, n) {
        fanya guess = x / 2; // Initial guess
        fanya tolerance = 0.0000000001; // Tolerance for convergence

        fanya calculateNthRoot = unda(x, n, guess, tolerance) {
            fanya nextGuess = ((n - 1) * guess + x / (guess ** (n - 1))) / n;
            fanya ipotolerance = abs(nextGuess - guess);
            kama (ipotolerance < tolerance) {rudisha nextGuess};
            rudisha calculateNthRoot(x, n, nextGuess, tolerance);
        }
        
        rudisha calculateNthRoot(x, n, guess, tolerance)
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

    //cos(x), calculates the cosine of an angle.
    cos = unda(x) {
        fanya result = 1; // Initialize the result
        fanya term = 1;

        kwa i ktk list(2,101,2) {
            term = (-term * x * x) / (i * (i - 1));
            result += term;
        }
        rudisha result;
    }

    //cosh(x), calculates the hyperbolic cosine of a number.
    cosh = unda(x) {
        fanya eToX = exp(x);
        fanya eToMinusX = exp(-x);
        rudisha(eToX + eToMinusX) / 2;
    }

    //exp(x), calculates the value of Euler's number raised to the power of a given number.
    exp = unda(n) {
        fanya result = 1;
        fanya term = 1;
       
        kwa i, v ktk list(1,23,1) {
            term = term*(n/v);
            result = result + term;
        }

        rudisha result;
    }


    //expm1(x), calculates the value of Euler's number raised to the power of a given number minus 1.
    expm1 = unda(x) {
        kama (x == -1) {
            rudisha -0.6321205588285577; // Handling the special case for -1
        } au kama (x == 0) {
            rudisha 0; // Handling the special case for 0
        } au kama (abs(x) < hisabati.EPSILON()) {
            rudisha x + 0.5 * x * x; // Approximation for very small x
        } sivyo {
            rudisha exp(x) - 1;
        }
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

        fanya sumOfSquares = reduce(values, exp, 0);

        // Calculate the square root of the sum of squares
        fanya result = sqrt(sumOfSquares);

        rudisha result;
    }

    //log(x), calculates the natural logarithm of a number.
    log = unda(x) {
        kama (x <= 0) {
            rudisha "NaN";
        }
        kama (x == 1) {
            rudisha 0;
        }
        kama (x < 0) {
            rudisha -log(-x);
        }
        fanya n = 1000; // Number of iterations
        fanya y = (x - 1) / (x + 1);
        fanya ySquared = y * y;
        fanya result = 0;
        kwa i ktk list(1,n+1,2) {
            result += (1 / i) * y**i;
        }
        rudisha 2 * result;
    }

    //log10(x), calculates the base 10 logarithm of a number.
    log10 = unda(x) {
        kama(x <= 0) {
            rudisha 0;
        }

        // Calculate natural logarithm and divide by the natural logarithm of 10
        rudisha log(x) / log(10.0);
    }

    //log1p(x), calculates the natural logarithm of 1 plus the given number.
    log1p = unda(x) {
        kama (x <= -1) {
            rudisha NaN; // Not a Number
        } au kama (abs(x) < hisabati.EPSILON()) {
            rudisha x - 0.5 * x * x; // Series expansion for small x
        } sivyo {
            rudisha log(1.0 + x);
        }
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
        fanya largest = 0;

        // Iterate through the numbers and update 'largest' kama a larger number is found
        kwa num ktk numbers {
            kama(num > largest) {
                largest = num;
            }
        }

        // rudisha the largest number (or 0 kama there are no parameters)
        rudisha largest;
    }

    //min(numbers), finds the minimum value in a list of numbers.
    min = unda(numbers) {
        kama(numbers.idadi() == 0) {
            rudisha 0;
        }

        fanya minVal = numbers[0];

        fanya i = 1;
        wakati(i < numbers.idadi()) {
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
    sin = unda(x) {
        fanya result = x; // Initialize the result with the angle
        fanya term = x;
        // Using Maclaurin series expansion for sine
        kwa i ktk list(3,101,2) {
            term = (-term * x * x) / (i * (i - 1));
            result += term;
        }
        rudisha result;
    }

    //sinh(x), calculates the hyperbolic sine of a number.
    sinh = unda(x) {
        // sinh(x) = (e^x - e^(-x)) / 2
        fanya eToX = exp(x);
        fanya eToMinusX = exp(-x);
        rudisha(eToX - eToMinusX) / 2;
    }

    //sqrt(x), calculates the square root of a number.
    sqrt = unda(x) {
        kama(x < 0) {
            rudisha 0;
        }
         kama(x >= 0) {
            rudisha root(x, 2);
        } sivyo {
            rudisha - root(-x, 2);
        }
    }

    //tan(x), calculates the tangent of an angle in radians.
    tan = unda(x) {
        fanya sineX = sin(x);
        fanya cosineX = sqrt(1 - sineX * sineX);

        kama(cosineX == 0) {
            rudisha 0;
        }

        rudisha sineX / cosineX;
    }

    //tanh(x), calculates the hyperbolic tangent of a number.
    tanh = unda(x) {
        fanya expX = exp(x);
        fanya expNegX = exp(-x);
        rudisha(expX - expNegX) / (expX + expNegX);
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
        rudisha sign(num)==-1;
    }

    //isInteger(num), checks if a number is an integer.
    isInteger = unda(num) {
        rudisha num == floor(num);
    }

    //getIntegerPart(num), gets the integer part of a number.
    getIntegerPart = unda(num) {
        // Handle negative numbers separately
        kama(isNegative(num)) {
            // For negative numbers, we subtract the absolute value of the fractional part from 1
            rudisha - (ceil(-num) - 1);
        } sivyo {
            // For positive numbers, we simply truncate the fractional part
            rudisha floor(num);
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
