namespace Application;

public static class AssemblyName
{
    public const string Name = "Application";
    public static readonly string Assembly = typeof(AssemblyName).Assembly.FullName!;
}