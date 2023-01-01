module TestAnswer

    using Test
    
    include("answer.jl")

    function main()
        @test Answer.divisor(6) == [1, 2, 3, 6]
        @test Answer.divisor(7) == [1, 7]
        @test Answer.q1(583305) == [1, 3, 5, 15, 37, 111, 185, 555, 1051, 3153, 5255, 15765, 38887, 116661, 194435, 583305] 

        @test Answer.isabundantnumber(6) == false
        @test Answer.isabundantnumber(12) == true
        @test Answer.q2() == 100000

        @test Answer.collatz(1) == [1, 4, 2, 1]
        @test Answer.collatz(2) == [2, 1]
        @test Answer.q3() == 20830
    end
end

if abspath(PROGRAM_FILE) == @__FILE__
    using .TestAnswer
    TestAnswer.main()
end
