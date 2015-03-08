#define CONSTYPE(T) ConsP_##T

#define MAKE_CONS(T) ;\
	type CONSTYPE(T) struct {;\
		car T;\
		cdr *CONSTYPE(T);\
	};\
   	func Cons_##T (t T) *CONSTYPE(T){;\
		return &CONSTYPE(T) {;\
			car: t,;\
		};\
	};\
	func (p *CONSTYPE(T)) Car() T {;\
		return p.car;\
	};\
	func (p *CONSTYPE(T)) Cdr() *CONSTYPE(T) {;\
		return p.cdr;\
	};\
	func (p *CONSTYPE(T)) Cons(n *CONSTYPE(T)) *CONSTYPE(T) {;\
		n.cdr = p;\
		return n;\
	};


