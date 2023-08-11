module Answer
function divisor(number)
    answer = []

    for i = 1:number
        if (number % i) == 0
            push!(answer, i)
        end
    end

    answer
end

function isabundantnumber(number)
    sum(divisor(number)) - number > number
end

function isdeficientnumber(number)
    sum(divisor(number)) - number < number
end

function isprime(number)
    for i = 2:floor(sqrt(number))
        if (number % i) == 0
            return false
        end
    end

    return true
end

function collatz(number)
    result = [number]

    while true
        if ((number % 2) == 0)
            number = number / 2
        else
            number = number * 3 + 1
        end

        push!(result, number)

        if number == 1
            return result
        end
    end
end

function q1(number)
    divisor(number)
end

function q2()
    number = 100000
    while number > 0
        if isabundantnumber(number)
            return number
        end

        number -= 1
    end
end

function q3()
    number = 1

    while true
        if size(collatz(number), 1) == 256
            return number
        end

        number += 1
    end
end


function q4(number)
    isprime(number)
end
end
