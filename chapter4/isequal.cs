namespace Generics
{
   private static void Main() {
      if(Compute<int>.IsEqual(2, 2)) {
            Console.WriteLine("2 isEqualTo 2");
         }
      if(!Compute<String>.IsEqual("A", "B")) {
            Console.WriteLine("A is_NOT_EqualTo B");
         }
   }
    public class Compute<T> {
        public static bool IsEqual(T Val1, T Val2) {
            return Val1.Equals(Val2);
        }
    }
}
