using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Net.Http;
using System.Reflection;
using System.Text;
using System.Text.RegularExpressions;

namespace ShellCommand
{
    public static class Util
    {
        public static (string, Dictionary<string, IEnumerable<string>>) ReadInput(Stream stream)
        {
            using (var reader = new StreamReader(stream))
            {
                var command = reader.ReadLine();

                var result = new Dictionary<string, IEnumerable<string>>();
                var pattern = new Regex("(?<key>.*): (?<value>.*)");

                string line;
                while ((line = reader.ReadLine()) != null)
                {
                    Log("input", line);
                    var match = pattern.Match(line);
                    var key = match.Groups["key"].Value;
                    var value = match.Groups["value"].Value;

                    if (!result.ContainsKey(key))
                    {
                        result[key] = new List<string>();
                    }

                    result[key] = result[key].Append(value);
                }

                return (command, result);
            }
        }

        public static void WriteOutput(Stream stream, IEnumerable<string> output, bool leaveOpen = false)
        {
            using (var writer = new StreamWriter(stream, Encoding.ASCII, -1, leaveOpen))
            {
                foreach (var line in output)
                {
                    writer.WriteLine(line);
                    Log("output", line);
                }
                writer.Flush();
            }
        }

        public static void Log(string tag, string message)
        {
            var dir = Path.GetDirectoryName(Assembly.GetEntryAssembly().Location);
            var file = Path.Join(dir, "log.txt");

            using (var writer = new StreamWriter(file, true))
            {
                var time = DateTime.Now.ToString();
                writer.WriteLine($"{time} - [{tag}] {message}");
            }
        }

        public static string GetDebugString(Dictionary<string, IEnumerable<string>> requests)
        {
            var sb = new StringBuilder();
            foreach (var request in requests)
            {
                sb.Append($"{request.Key}=[{string.Join(",", request.Value)}], ");
            }

            return sb.ToString();
        }
    }
}
