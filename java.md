```java
public class Lox {
  public static void main(String[] args) throws IOException {
    if (args.length > 1) {
      System.out.println("Usage: jlox [script]");
      System.exit(64);
    } else if (args.length == 1) {
      runFile(args[0]);
    } else {
      runPrompt();
    }
  }
}

```

```java
private static void runFile(String path) throws IOException {
byte[] bytes = Files.readAllBytes(Paths.get(path));
run(new String(bytes, Charset.defaultCharset()));
}
```

```java

```

```java

```

```java

```