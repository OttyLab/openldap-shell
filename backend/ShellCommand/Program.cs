using ShellCommand.Db;
using System;

namespace ShellCommand
{
    class Program
    {
        static void Main(string[] args)
        {
            using (var db = new LocalFileDb())
            {
                (var command, var requests) = Util.ReadInput(Console.OpenStandardInput());

                switch (command)
                {
                    case "ADD":
                        Util.Log("ADD", Util.GetDebugString(requests));
                        db.Add(requests);
                        break;
                    case "SEARCH":
                        Util.Log("SEARCH", Util.GetDebugString(requests));
                        Util.WriteOutput(Console.OpenStandardOutput(), db.Search(requests));
                        break;
                    default:
                        Util.Log(command, Util.GetDebugString(requests));
                        throw new NotImplementedException($"{command} is not supported");
                }
            }
        }
    }
}
