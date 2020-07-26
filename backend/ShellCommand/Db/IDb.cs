using System;
using System.Collections.Generic;

namespace ShellCommand.Db
{
    public interface IDb : IDisposable
    {
        bool Add(Dictionary<string, IEnumerable<string>> request);
        IEnumerable<string> Search(Dictionary<string, IEnumerable<string>> request);
        public IEnumerable<string> Compare(Dictionary<string, IEnumerable<string>> requests);
    }
}
