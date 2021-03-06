package main

import (
    "fmt"
    "os"
    "time"
)

func validateTree(n *tnode) {
    parent := n.ntype
    if parent == T_STARTSWITH || parent == T_ICSTARTSWITH ||
            parent == T_ENDSWITH || parent == T_ICENDSWITH ||
            parent == T_CONTAINS || parent == T_ICCONTAINS {
        if right(n).ntype != T_STRING {
            fail()
        }
    } else if parent == T_EQ || parent == T_ICEQ || parent == T_NEQ ||
            parent == T_ICNEQ || parent == T_LT || parent == T_LTE ||
            parent == T_GT || parent == T_GTE {
        l := left(n).ntype
        r := right(n).ntype
        if l == T_NAME || l == T_PATH || l == T_CONTENT {
            if r != T_STRING {
                fail()
            }

            if parent == T_LT || parent == T_LTE ||
                    parent == T_GT || parent == T_GTE {
                fail()
            }
        } else if l == T_SIZE {
            if r != T_INTEGER {
                fail()
            }
        } else if l == T_MODIFIED {
            if r != T_STRING {
                fail()
            }

            _, derr := time.Parse(DATE_FORMAT, right(n).sval)
            _, terr := time.Parse(TIMESTAMP_FORMAT, right(n).sval)
            if derr != nil && terr != nil {
                fmt.Println("invalid expression - expected modified time to be in one of the following formats: " +
                    "\"" + TIMESTAMP_FORMAT + "\" or \"" + DATE_FORMAT + "\"")
                os.Exit(1)
            }
        }
    }

    if len(n.children) > 0 {
        validateTree(left(n))
        if len(n.children) > 1 {
            validateTree(right(n))
        }
    }
}

func fail() {
    fmt.Println("invalid expression")
    os.Exit(1)
}
