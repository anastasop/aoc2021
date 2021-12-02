/forward/ { x += $2; y += aim * $2 }
/up/ { aim -= $2 }
/down/ { aim += $2 }
END { print x * y }
