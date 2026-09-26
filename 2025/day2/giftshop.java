import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;

public class giftshop {
    public static void main(String[] args) throws IOException {
        long totalInvalid = 0;
        Path input = Path.of("input.txt");

        String content = Files.readString(input);
        
        String[] parts = content.split(",");
        for (String part : parts) {
            String[] numbers = part.split("-");
            long start = Long.valueOf(numbers[0]);
            long end = Long.valueOf(numbers[1]);
            String startS = numbers[0];
            String endS = numbers[1];
            for(long i = start; i < end + 1; i++) {
                long number = i;
                String numberS = String.valueOf(number);
                if(numberS.length() % 2 == 0) {
                    int mid = numberS.length() / 2;
                    String[] partsNumber = { numberS.substring(0, mid), numberS.substring(mid) };
                    if(partsNumber[0].equals(partsNumber[1])) {
                        totalInvalid+=number;
                    }
                }
            }
        }
        System.out.println(String.valueOf(totalInvalid));
    }      
}
