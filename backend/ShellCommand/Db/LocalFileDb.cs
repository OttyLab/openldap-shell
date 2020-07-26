using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Reflection;
using System.Text;
using System.Text.Json;
using System.Text.RegularExpressions;
using System.Threading.Tasks;

namespace ShellCommand.Db
{
    using Contents = Dictionary<string, IEnumerable<string>>;
    using Schema = Dictionary<string, Dictionary<string, IEnumerable<string>>>;

    public class LocalFileDb : IDb
    {
        private Stream Stream { get; }

        public LocalFileDb(Stream stream = default)
        {
            if (stream != default)
            {
                Stream = stream;
                return;
            }

            var dir = Path.GetDirectoryName(Assembly.GetEntryAssembly().Location);
            var file = Path.Join(dir, "db.json");

            while (true)
            {
                try
                {
                    Stream = new FileStream(file, FileMode.OpenOrCreate, FileAccess.ReadWrite, FileShare.None);
                    break;
                }
                catch (IOException)
                {
                    Console.WriteLine("waiting");
                    Task.Delay(10000).Wait();
                }
            }
        }

        public void Dispose()
        {
            Stream.Close();
        }

        public bool Add(Contents request)
        {
            var storage = ReadDb();

            string dn = "";
            var contents = new Contents();

            foreach ((var key, var value) in request)
            {
                if (key == "dn")
                {
                    dn = value.ElementAt(0);
                }

                if (key == "msgid" || key == "suffix")
                {
                    //Do nothing
                    continue;
                }
                else
                {
                    contents[key] = value;
                }
            }

            storage[dn] = contents;

            var serialized = JsonSerializer.Serialize(storage);
            Stream.SetLength(0);
            Stream.Write(Encoding.UTF8.GetBytes(serialized));

            return true;
        }

        public IEnumerable<string> Search(Contents requests)
        {
            var storage = ReadDb();

            var scope = int.Parse(requests["scope"].ElementAt(0));
            var sizeLimit = int.Parse(requests["sizelimit"].ElementAt(0));

            var pattern = $"^([^=]+=[^=,]+,){{{scope}}}{requests["base"].ElementAt(0)}";
            var entries = SearchInternal(storage, pattern, requests["filter"].ElementAt(0));
            var result = new List<string>();

            var cnt = 1;
            foreach(var entry in entries)
            {
                if (sizeLimit > 0 && cnt++ > sizeLimit)
                {
                    break;
                }

                foreach(var item in entry)
                {
                    foreach(var value in item.Value)
                    {
                        result.Add($"{item.Key}: {value}");
                    }
                }

                result.Add("");
            }

            result.Add("RESULT");
            result.Add("code: 0");

            return result;
        }

        public IEnumerable<string> Compare(Contents requests)
        {
            var storage = ReadDb();
            var result = new List<string>();

            var dn = requests["dn"].ElementAt(0);
            if (!storage.ContainsKey(dn))
            {
                result.Add("RESULT");
                result.Add("code: 34");
                return result;
            }

            var target = storage[dn];

            foreach (var key in requests.Keys)
            {
                if (key == "dn" || key == "msgid" || key == "suffix")
                {
                    continue;
                }

                if (!target.ContainsKey(key))
                {
                    result.Add("RESULT");
                    result.Add("code: 32");
                    return result;
                }

                if (!target[key].Contains(requests[key].ElementAt(0)))
                {
                    result.Add("RESULT");
                    result.Add("code: 5");
                    return result;
                }
            }

            result.Add("RESULT");
            result.Add("code: 6");

            return result;
        }

        private Schema ReadDb()
        {
            if (Stream.Length == 0)
            {
                return new Schema();
            }

            var buffer = new byte[Stream.Length];
            Stream.Read(buffer);
            var text = Encoding.UTF8.GetString(buffer);
            return JsonSerializer.Deserialize<Schema>(text);
        }

        private List<Contents> SearchInternal(Schema storage, string pattern, string filter)
        {
            var results = new List<Contents>();
            foreach(var item in storage)
            {
                if (Regex.IsMatch(item.Key, pattern) && DoesPassFilter(item.Value, filter))
                {
                    results.Add(item.Value);
                }
            }

            return results;
        }

        private bool DoesPassFilter(Contents contents, string filter)
        {
            var pattern = "\\((?<key>[^&]+?)=(?<value>.+?)\\)";
            var matches = Regex.Matches(filter, pattern);

            // TODO: think of "|"
            foreach (Match match in matches)
            {
                var key = match.Groups["key"].Value;
                var value = match.Groups["value"].Value;

                if (!contents.ContainsKey(key))
                {
                    return false;
                }

                if (value == "*")
                {
                    continue;
                }

                if (!Enumerable.Contains(contents[key], value))
                {
                    return false;
                }
            }

            return true;
        }
    }
}
