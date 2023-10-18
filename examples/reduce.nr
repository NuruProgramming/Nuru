fanya reduce = unda(iterator, callback, initialValue) {
  fanya accumulator = initialValue;

  kwa thamani ktk iterator {
    accumulator = callback(accumulator, thamani);
  }

  rudisha accumulator;
}

fanya list = [1,2,3,4,5];
fanya employees = [{"salary":120},{"salary":135},{"salary":140}]

fanya sum = unda(acc,value){
  rudisha acc + value;
}

fanya mul = unda(acc,value){
  rudisha acc * value;
}

fanya sumSalo = unda(acc,value){
  rudisha acc + value["salary"];
}

fanya sumSaloWithTax = unda(acc,value){
  rudisha acc + (value["salary"] * (1-0.34));
}

andika(reduce(list,sum,0))
andika(reduce(list,mul,1))

andika(reduce(employees,sumSalo,0))
andika(reduce(employees,sumSaloWithTax,0))