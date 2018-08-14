arr = []
arr1
for i in range(5):
    func =lambda x,i=i:i+x
    arr1.append(func)
    

for f in arr :
    print(f(5))
