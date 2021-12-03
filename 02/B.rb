
$x = $y = $a = 0 

def forward(n)
  $x += n
  $y += n * $a
end

def down(n)
  $a += n
end

def up(n)
  $a -= n
end

eval(File.read('input.txt'))
puts $x * $y

