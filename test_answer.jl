module TestAnswer
using Test

include("answer.jl")

function main()
    @testset "Divisor" begin
        @test Answer.divisor(6) == [1, 2, 3, 6]
        @test Answer.divisor(7) == [1, 7]
    end

    @testset "Sum of divisior" begin
        @test Answer.isabundantnumber(12) == true
        @test Answer.isabundantnumber(6) == false
    end

    @testset "Prime" begin
        @test Answer.isprime(3) == true
        @test Answer.isprime(4) == false
    end

    @testset "Collatz" begin
        @test Answer.collatz(1) == [1, 4, 2, 1]
        @test Answer.collatz(2) == [2, 1]
    end

    @testset "Kaitou" begin
        @test Answer.q1(583305) == [
            1,
            3,
            5,
            15,
            37,
            111,
            185,
            555,
            1051,
            3153,
            5255,
            15765,
            38887,
            116661,
            194435,
            583305,
        ]
        @test Answer.q2() == 100000
        @test Answer.q3() == 20830
        @test Answer.q4(1356361) == false
    end
end
end

if abspath(PROGRAM_FILE) == @__FILE__
    using .TestAnswer

    TestAnswer.main()
end
