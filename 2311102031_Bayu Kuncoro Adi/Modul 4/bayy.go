procedure mencariFaktorial(n: integer) -> integer
    if n == 0 or n == 1 then
        return 1
    else
        return n * mencariFaktorial(n - 1)

function permutation(n: integer, r: integer) -> integer
    return mencariFaktorial(n) / mencariFaktorial(n - r)

function combination(n: integer, r: integer) -> integer
    return mencariFaktorial(n) / (mencariFaktorial(r) * mencariFaktorial(n - r))

procedure main()
    input n1, r1, n2, r2  // Masukan 4 bilangan asli
    P1 = permutation(n1, r1)
    C1 = combination(n1, r1)
    P2 = permutation(n2, r2)
    C2 = combination(n2, r2)
    
    // Output
    print("P(", n1, ",", r1, ") = ", P1)
    print("C(", n1, ",", r1, ") = ", C1)
    print("P(", n2, ",", r2, ") = ", P2)
    print("C(", n2, ",", r2, ") = ", C2)









procedure zoomIn(w: integer, h: integer, skala: integer) -> (integer, integer)
    // Memperbesar ukuran w dan h sebesar skala
    return w * skala, h * skala

procedure zoomOut(w: integer, h: integer, skala: integer) -> (integer, integer)
    // Memperkecil ukuran w dan h sebesar skala
    return w / skala, h / skala

procedure main()
    // Input
    input w, h  // ukuran panjang dan lebar gambar
    input s, skala  // simbol operasi dan skala
    
    if s == '+' then
        w, h = zoomIn(w, h, skala)
    else if s == '-' then
        w, h = zoomOut(w, h, skala)
    
    // Output
    print(w, h)