class Solution {
    public:
        int maximum69Number (int num) {
            int cur = 0;
            int numC = num;
            int indFirstSix = INT_MIN;
            while(numC!=0) {
                if(numC%10 == 6){
                    indFirstSix = cur;
                }
                numC /= 10;
                cur++;
            }
    
            return indFirstSix == -1 ? num : num + 3 * (int)pow(10, indFirstSix);
        }
    };