tumia "hesabu"

andika("abs: ",hesabu.abs(-42));
andika("acos: ",hesabu.acos(0.5));
andika("acosh: ",hesabu.acosh(2));
andika("asin: ", hesabu.asin(0.5));
andika("asinh: ",hesabu.asinh(2));
andika("atan: ",hesabu.atan(1));
andika("atan2: ",hesabu.atan2(1, 1));
andika("atanh: ",hesabu.atanh(0.5));
andika("cbrt: ",hesabu.cbrt(8));
andika("root: ",hesabu.root(27, 3));
andika("ceil: ",hesabu.ceil(4.3));
andika("cos: ",hesabu.cos(5));
andika("cosh: ",hesabu.cosh(5));
andika("exp: ",hesabu.exp(2));
andika("expm1: ",hesabu.expm1(1));
andika("floor: ",hesabu.floor(4.7));
andika("hypot: ",hesabu.hypot([3, 4]));
andika("log: ",hesabu.log(2));
andika("log10: ",hesabu.log10(100));
andika("log1p: ",hesabu.log1p(1));
andika("log2: ",hesabu.log2(8));
andika("max: ",hesabu.max([4, 2, 9, 5]));
andika("min: ",hesabu.min([4, 2, 9, 5]));
andika("round: ",hesabu.round(4.6, "rpi"));
andika("sign: ",hesabu.sign(-5));
andika("sin: ",hesabu.sin(1));
andika("sinh: ",hesabu.sinh(0));
andika("sqrt: ",hesabu.sqrt(4));
andika("tan: ",hesabu.tan(1));
andika("tanh: ",hesabu.tanh(0));
andika("factorial: ",hesabu.factorial(5));
andika("isNegative: ",hesabu.isNegative(-5));
andika("isInteger: ",hesabu.isInteger(4.5));
andika("getIntegerPart: ",hesabu.getIntegerPart(4.5));
andika("list: ",hesabu.list(1, 5, 1));
fanya callback = unda(accumulator, currentValue){
    rudisha accumulator + currentValue;
}
andika("reduce: ",hesabu.reduce([1, 2, 3, 4],callback,0) );
