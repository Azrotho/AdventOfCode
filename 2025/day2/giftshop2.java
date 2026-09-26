import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.List;

public class giftshop2 {
    public static void main(String[] args) throws IOException {
        long totalInvalid = 0;
        Path input = Path.of("input.txt");

        String content = Files.readString(input);
        int actual = 0;
        
        String[] parts = content.split(",");
        for (String part : parts) {
            actual += 1;
            String[] numbers = part.split("-");
            long start = Long.valueOf(numbers[0]);
            long end = Long.valueOf(numbers[1]);
            String startS = numbers[0];
            String endS = numbers[1];
            for(long i = start; i <= end; i++) {
                String numberS = String.valueOf(i);
                int len = numberS.length();
                for (int n = 2; n <= len; n++) {
                    if (len % n == 0) {
                        if (checkSplit(numberS, n)) {
                            totalInvalid += i;
                            break;
                        }
                    }
                }
            }
            System.out.printf("[%d/%d]\n", actual, parts.length);
        }
        System.out.println(String.valueOf(totalInvalid));
    }

    public static List<Long> printDivisors(long n) {
        List<Long> div = new ArrayList<>();
        for (long i = 1; i <= Math.sqrt(n); i++) {
            if (n % i == 0) { 
                //System.out.print(i + " ");
                div.add(i);
                if (i != n / i) {
                    //System.out.print(n / i + " ");
                    div.add(n / i);
                }
            }
        }
        return div;
    }

public static boolean checkSplit(String text, int parts) {
    int len = text.length();
    int partSize = len / parts;

    String pattern = text.substring(0, partSize);

    for (int i = 1; i < parts; i++) {
        int start = i * partSize;
        int end = start + partSize;
        
        if (!text.substring(start, end).equals(pattern)) {
            return false;
        }
    }
    return true;
}
}
