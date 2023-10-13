subconj([], _).      % Un conjunto vacío es subconjunto de cualquier conjunto.
subconj([H|T], S) :- member(H, S), subconj(T, S).
