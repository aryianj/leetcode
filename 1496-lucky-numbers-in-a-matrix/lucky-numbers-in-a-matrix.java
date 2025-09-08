class Solution {
    public List<Integer> luckyNumbers(int[][] matrix) {
        List<Integer> minList = new ArrayList<>();
        List<Integer> maxList = new ArrayList<>();
        
        for (int x = 0; x < matrix.length; x++) {
            List<Integer> list = new ArrayList<>();
            for (int y = 0; y < matrix[0].length; y++) {
                list.add(matrix[x][y]);
            }
            minList.add(Collections.min(list));
        }

        for (int y = 0; y < matrix[0].length; y++) {
            List<Integer> list = new ArrayList<>();
            for (int x = 0; x < matrix.length; x++) {
                list.add(matrix[x][y]);
            }
            maxList.add(Collections.max(list));
        }

        minList.retainAll(maxList);
        return minList;
    }
}