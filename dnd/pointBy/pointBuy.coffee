budget = Number(process.argv[2])
minStat = Number(process.argv[3]) or 3
maxStat = Number(process.argv[4]) or 18

unless budget and minStat and maxStat
    console.error "Invalid point buy number"
    process.exit 1


pbCost =
    3:  -5
    4:  -4
    5:  -3
    6:  -2
    7:  -1
    8:  0
    9:  1
    10: 2
    11: 3
    12: 4
    13: 5
    14: 7
    15: 9
    16: 11
    17: 14
    18: 17


RES = {}

isFinalStat = (char, budget) ->
    cost = 0
    cost += pbCost[x] for x in char
    rest = budget - cost
    improveCost = pbCost[char[0]+1] - pbCost[char[0]]

    return false if rest < 0                   # not enought
    return false if rest >= improveCost        # possible to enchance
    return true


statValues = [minStat..maxStat]

console.log "Budget: #{budget}, Stat Range: #{minStat}-#{maxStat}"

for str in statValues
    for dex in statValues
        for con in statValues
            for int in statValues
                for wis in statValues
                    for cha in statValues

                        char = [str, dex, con, int, wis, cha]
                        char.sort (a,b) -> a > b
                        continue unless isFinalStat char, budget
                        char = char.reverse().map (o) -> return if o < 10 then "0#{o}" else "#{o}"

                        RES[ char.join(', ') ] = 1


console.log Object.keys(RES).sort().reverse().join("\n")

# xxx = [ 15, 15, 13, 13, 10, 9 ]
# xxx.sort (a,b) -> a > b
# console.log isFinalStat xxx, 32
