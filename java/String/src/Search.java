public class Search {
    public static void main(String[] args) {
        String source = "hello wcm";
        String niddle = "wcm";

        int lastIndex = source.lastIndexOf(niddle);

        if(lastIndex == -1) {
            System.out.println("not found");
        } else {
            System.out.println("found");
        }

    }
}
