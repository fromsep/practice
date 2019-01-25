public class Compare {

    public static void main(String[] args) {
        String strA = "hello wcm";
        String strB = "wcm";

        Object obj = strA;

        System.out.println(strA.compareTo(strB));
        System.out.println(strA.compareToIgnoreCase(strB));
        System.out.println(strA.compareTo(obj.toString()));
    }
}
