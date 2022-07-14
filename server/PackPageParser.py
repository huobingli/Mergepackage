import sys

if __name__ == '__main__':
    
    if len(sys.argv) != 2:
        exit()    
    
    ishq = sys.argv[1]

    if ishq == "1":
        print("hq")
    else:
        print("jy")

    # /str = "[{value:\"test444\", label:\"tesee\"}, {value:\"1233312\", label:\"66666\"}]"
    str = "[{value:\"release-8.70.50.111\",label:\"release-8.70.50.111\"},{value:\"release-8.70.50.112\",label:\"release-8.70.50.112\"},{value:\"release-8.70.50.113\",label:\"release-8.70.50.113\"}]"
    print(str)
    