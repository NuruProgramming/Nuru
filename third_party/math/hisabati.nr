pakeji hisabati{
    andaa = unda() { // the andaa function is mandatory even kama its empty
        // π (Pi): The ratio of the circumference of a circle to its diameter.
        @.PI = 3.141592653589793;

        // e (Euler's Number): The base of the natural logarithm and used in various mathematical contexts.
        @.e = 2.718281828459045;

        // φ (Phi, Golden Ratio): A mathematical fanyaant with applications in nature, art, and architecture.
        @.phi = 1.618033988749895;

        // Natural logarithm of 10.
        @.ln10 =  2.302585092994046;

        // Natural logarithm of 2.
        @.ln2 = 0.6931471805599453;

        // Base 10 logarithm of Euler's number (e).
        @.log10e =  0.4342944819032518;

        // Base 2 logarithm of Euler's number (e).
        @.log2e =  1.4426950408889634;

        // √1/2 (equivalent to 1 / sqrt(2)).
        @.sqrt1_2 = 0.7071067811865476;

        // √2 (Square Root of 2): Used in geometry and mathematical proofs.
        @.sqrt2 = 1.414213562373095;

        // √3 (Square Root of 3): Used in geometry, trigonometry, and mathematics.
        @.sqrt3 = 1.732050807568877;

        // √5 (Square Root of 5): Used in various mathematical contexts.
        @.sqrt5 = 2.236067977499790;

        //@.EPSILON = 2.220446049250313e-16;

    }

    
    // Methods

    abs = unda(namba){
        kama(namba < 0){
            rudisha - 1 * namba;
        }

        rudisha namba;
    }

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

    acosh = unda(x) {
        kama(x < 1) {
            rudisha 0;
        }

        rudisha @.log(x + @.sqrt(x * x - 1));
    }

    // method to calculate arcsin using Taylor series
    arcsin = unda(x) {
        kama(x < -1 || x > 1) {
            rudisha 0;
        }

        fanya maxIterations = 50; // Maximum number of iterations
        fanya result = 0;

        fanya n = 0;
        wakati(n < maxIterations) {
            fanya numerator = @.pow(-1, n) * @.pow(x, 2 * n + 1);
            fanya denominator = @.factorial(2 * n + 1); // You'll need to implement a factorial function

            result += numerator / denominator;
            n++;
        }

        rudisha result;
    }

    arsinh = unda(x) {
        // Calculate arsinh using the formula: arsinh(x) = ln(x + sqrt(x^2 + 1))
        kama(x >= 0) {
            rudisha @.log(x + @.sqrt(x * x + 1));
        } sivyo {
            // For negative values, arsinh(x) = -arsinh(-x)
            rudisha - @.log(-x + @.sqrt(x * x + 1));
        }
    }

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

    atanh = unda(x) {
        kama(x < -1 || x > 1) {
            rudisha 0;
        }
        rudisha 0.5 * @.log((1 + x) / (1 - x));
    }

    // Cube Root
    cbrt = unda(x) {
        kama(x >= 0) {
            rudisha @.root(x, 3);
        } sivyo {
            rudisha - @.root(-x, 3);
        }
    }

    // Helper method to calculate the nth root using the Newton-Raphson method
    root = unda(x, n) {
        fanya guess = x / 2; // Initial guess
        fanya tolerance = 1e-10; // Tolerance for convergence

        wakati(true) {
            fanya nextGuess = ((n - 1) * guess + x / @.pow(guess, n - 1)) / n;
            kama(@.abs(nextGuess - guess) < tolerance) {
                rudisha nextGuess;
            }
            guess = nextGuess;
        }
    }

    // Rounds up to the smallest integer greater than or equal to a given number
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
            fanya denominator = factorial(2 * n);

            // Add the nth term to the result
            result += numerator / denominator;

            n++;
        }

        rudisha result;
    }

    cosh = unda(x) {
        fanya eToX = @.exp(x);
        fanya eToMinusX = @.exp(-x);
        rudisha(eToX + eToMinusX) / 2;
    }

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

    floor = unda(x) {
        kama(x >= 0) {
            rudisha x - (x % 1);
        } sivyo {
            rudisha x - (1 + x % 1);
        }
    }

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

    log = unda(x) {
        kama(x <= 0) {
            rudisha 0;
        }

        fanya approx = 0;
        fanya n = 50;
        fanya i = 1;

        wakati(i <= n) {
            approx += (1 / i) * @.pow((x - 1) / (x + 1), 2 * i - 1);
            i++;
        }


        rudisha 2 * approx;
    }

    // method to calculate the base 10 logarithm
    log10 = unda(x) {
        kama(x <= 0) {
            rudisha 0;
        }

        // Calculate natural logarithm and divide by the natural logarithm of 10
        rudisha this.log(x) / this.log(10);
    }

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

    sign = unda(x) {
        kama(x == 0 || x == -0) {
            rudisha x;
        } au kama(x > 0) {
            rudisha 1;
        } sivyo {
            rudisha - 1;
        }
    }

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

    sinh = unda(x) {
        // sinh(x) = (e^x - e^(-x)) / 2
        fanya eToX = @.exp(x);
        fanya eToMinusX = @.exp(-x);
        rudisha(eToX - eToMinusX) / 2;
    }

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

    sine = unda(x) {
        fanya tolerance = @.e-15; // Tolerance for convergence
        fanya result = x;
        fanya term = x;
        fanya n = 1;

        wakati(@.abs(term) > tolerance) {
            term *= (-x * x) / ((2 * n + 1) * (2 * n));
            result += term;
            n++;
        } 

        rudisha result;
    }

    // Calculate the tangent of a number in radians
    tangent = unda(x) {
        fanya sineX = @.sine(x);
        fanya cosineX = @.sqrt(1 - sineX * sineX);

        kama(cosineX == 0) {
            rudisha 0;
        }

        rudisha sineX / cosineX;
    }

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
    // Function to check kama a number is negative zero
    isNegativeZero = unda(num) {
        rudisha num == 0 && 1 / num == -Infinity;
    }
    // Helper function to calculate the factorial of a number
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

    // Helper function to calculate pow using Taylor series
    pow = unda(base, exponent) {
        fanya result = 1;
        fanya i = 0;
        wakati(i < exponent) {
            result *= base;
            i++;
        }
        rudisha result;
    }

    // Helper function to check kama a number is negative
    isNegative = unda(num) {
        rudisha num < 0;
    }

    isInteger = unda(num) {
        rudisha num == @.floor(num);
    }

    // method to get the integer part of a number
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

    // create a list of numbers
    list = unda(first, last, interval){
        fanya list = [first];
        fanya i = first + interval;
        wakati(i < last){
            list.sukuma(i);
            i += interval;
        }
        rudisha list;
    }

    // Maths functions
    // square a number
    square = unda(n, i, j){
        fanya kati = (i + j) / 2;
        fanya mul = kati * kati;
        fanya abs_diff = abs(mul - n);

        kama(mul == n || abs_diff < 0.00001){
            rudisha kati;
        }au kama(mul < n){
            rudisha square(n, kati, j);
        }au{
            rudisha square(n, i, kati);
        }
    }

    // find the square root of a number
    sqrt = unda(namba){
        kwa i ktk list(0, namba, 1) {
            kama((i * i) == namba){
            rudisha i;
            }au kama((i * i) > namba){
            rudisha square(namba, i - 1, i);
            }
        }
    }

    reduce = unda(iterator, callback, initialValue) {
        fanya accumulator = initialValue;

        kwa thamani ktk iterator {
            accumulator = callback(accumulator, thamani);
        }

        rudisha accumulator;
    }
}
