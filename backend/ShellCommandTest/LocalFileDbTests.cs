using NuGet.Frameworks;
using NUnit.Framework;
using ShellCommand.Db;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Text.Json;

namespace ShellCommandTest
{
    using Contents = Dictionary<string, IEnumerable<string>>;
    using Schema = Dictionary<string, Dictionary<string, IEnumerable<string>>>;

    public class LocalFileDbTests
    {
        private IDb db;
        private Stream stream;

        [SetUp]
        public void Setup()
        {
            stream = new MemoryStream();
            db = new LocalFileDb(stream);
        }

        [TearDown]
        public void TearDown()
        {
            db.Dispose();
        }

        [Test]
        public void TestAdd()
        {
            var requests = new Contents()
            {
                {"dn", new List<string> { "dn0" }},
                {"msgid", new List<string> { "1" }},
                {"suffix", new List<string> { "suffix" }},
                {"a", new List<string> { "a0" }},
                {"b", new List<string> { "b0", "b1" }},
            };

            var result = db.Add(requests);
            Assert.IsTrue(result);

            stream.Position = 0;
            var buffer = new byte[stream.Length];
            stream.Read(buffer);
            var actual = Encoding.UTF8.GetString(buffer);
            var expected = "{\"dn0\":{\"dn\":[\"dn0\"],\"a\":[\"a0\"],\"b\":[\"b0\",\"b1\"]}}";
            Assert.AreEqual(expected, actual);

            stream.Position = 0;
            requests = new Contents()
            {
                {"dn", new List<string> { "dn1" }},
                {"c", new List<string> { "c0", "c1", "c2" }},
            };

            result = db.Add(requests);
            Assert.IsTrue(result);

            stream.Position = 0;
            buffer = new byte[stream.Length];
            stream.Read(buffer);
            actual = Encoding.UTF8.GetString(buffer);
            expected = "{\"dn0\":{\"dn\":[\"dn0\"],\"a\":[\"a0\"],\"b\":[\"b0\",\"b1\"]},\"dn1\":{\"dn\":[\"dn1\"],\"c\":[\"c0\",\"c1\",\"c2\"]}}";
            Assert.AreEqual(expected, actual);
        }

        [Test]
        public void TestSearch()
        {
            var source = new Schema
            {
                { "dc=example,dc=com", new Contents {
                    { "dn", new List<string>(){"dc=example,dc=com"} },
                    { "objectClass", new List<string>(){ "organization", "dcObject" } },
                    { "o", new List<string>(){ "Example Organization" } },
                    { "dc", new List<string>(){ "example" } },
                } },
                { "ou=Employee,dc=example,dc=com", new Contents {
                    { "dn", new List<string>(){"ou=Employee,dc=example,dc=com"} },
                    { "objectClass", new List<string>(){ "organizationalUnit" } },
                    { "ou", new List<string>(){ "Employee" } },
                } },
                { "cn=taro.yamada,ou=Employee,dc=example,dc=com", new Contents {
                    { "dn", new List<string>(){"cn=taro.yamada,ou=Employee,dc=example,dc=com"} },
                    { "objectClass", new List<string>(){ "inetOrgPerson", "posixAccount" } },
                    { "cn", new List<string>(){ "Taro Yamada" } },
                    { "sn", new List<string>(){ "Yamada" } },
                    { "uid", new List<string>(){ "taro.yamada" } },
                    { "userPassword", new List<string>(){ "{SSHA}46mXVpqhvxX3mF+yAzawY47d6ldDwwAs" } },
                    { "uidNumber", new List<string>(){ "3001" } },
                    { "gidNumber", new List<string>(){ "3000" } },
                    { "loginShell", new List<string>(){ "/bin/bash" } },
                    { "homeDirectory", new List<string>(){ "/home/taro.yamada" } },
                } },
            };

            var json = JsonSerializer.Serialize(source);
            stream.Write(Encoding.UTF8.GetBytes(json));

            stream.Position = 0;

            var requests = new Contents()
            {
                { "msgid", new List<string>(){ "1" } },
                { "suffix", new List<string>(){ "dc=example,dc=com" } },
                { "base", new List<string>(){ "dc=example,dc=com" } },
                { "scope", new List<string>(){ "0" } },
                { "deref", new List<string>(){ "0" } },
                { "sizelimit", new List<string>(){ "-1" } },
                { "timelimit", new List<string>(){ "-1" } },
                { "filter", new List<string>(){ "" } },
                { "attrsonlyr", new List<string>(){ "0" } },
                { "attrs", new List<string>(){ "all" } },
            };

            var actual = db.Search(requests);
            Assert.AreEqual("dn: dc=example,dc=com", actual.ElementAt(0));
            Assert.AreEqual("objectClass: organization", actual.ElementAt(1));
            Assert.AreEqual("objectClass: dcObject", actual.ElementAt(2));
            Assert.AreEqual("o: Example Organization", actual.ElementAt(3));
            Assert.AreEqual("dc: example", actual.ElementAt(4));
            Assert.AreEqual("", actual.ElementAt(5));
            Assert.AreEqual("RESULT", actual.ElementAt(6));
            Assert.AreEqual("code: 0", actual.ElementAt(7));
        }

        [Test]
        public void TestSearchWithFilter()
        {
            var requests = new Contents()
            {
                {"dn", new List<string> { "cn=admin,ou=Employee,dc=example,dc=com" }},
                {"objectClass", new List<string> { "inetOrgPerson", "posixAccount"}},
                {"cn", new List<string> { "admin" }},
                {"uid", new List<string> { "admin" }},
            };

            db.Add(requests);

            stream.Position = 0;

            requests = new Contents()
            {
                {"dn", new List<string> { "cn=taro.yamada,ou=Employee,dc=example,dc=com" }},
                {"objectClass", new List<string> { "inetOrgPerson", "posixAccount"}},
                {"cn", new List<string> { "Taro Yamada" }},
                {"uid", new List<string> { "taro.yamada" }},
                {"uidNumber", new List<string> { "3001" }},
            };

            db.Add(requests);

            stream.Position = 0;

            requests = new Contents()
            {
                { "msgid", new List<string>(){ "1" } },
                { "suffix", new List<string>(){ "dc=example,dc=com" } },
                { "base", new List<string>(){ "dc=example,dc=com" } },
                { "scope", new List<string>(){ "2" } },
                { "deref", new List<string>(){ "0" } },
                { "sizelimit", new List<string>(){ "-1" } },
                { "timelimit", new List<string>(){ "-1" } },
                { "filter", new List<string>(){ "(&(objectClass=*)(uid=taro.yamada)" } },
                { "attrsonlyr", new List<string>(){ "0" } },
                { "attrs", new List<string>(){ "uid" } },
            };

            var actual = db.Search(requests);
            Assert.AreEqual("dn: cn=taro.yamada,ou=Employee,dc=example,dc=com", actual.ElementAt(0));
            Assert.AreEqual("objectClass: inetOrgPerson", actual.ElementAt(1));
            Assert.AreEqual("objectClass: posixAccount", actual.ElementAt(2));
            Assert.AreEqual("cn: Taro Yamada", actual.ElementAt(3));
            Assert.AreEqual("uid: taro.yamada", actual.ElementAt(4));
            Assert.AreEqual("uidNumber: 3001", actual.ElementAt(5));

            stream.Position = 0;

            requests = new Contents()
            {
                { "msgid", new List<string>(){ "1" } },
                { "suffix", new List<string>(){ "dc=example,dc=com" } },
                { "base", new List<string>(){ "dc=example,dc=com" } },
                { "scope", new List<string>(){ "2" } },
                { "deref", new List<string>(){ "0" } },
                { "sizelimit", new List<string>(){ "-1" } },
                { "timelimit", new List<string>(){ "-1" } },
                { "filter", new List<string>(){ " (&(uid=taro.yamada)(objectClass=posixAccount)(&(uidNumber=*)(!(uidNumber=0))))" } },
                { "attrsonlyr", new List<string>(){ "0" } },
                { "attrs", new List<string>(){ "uid" } },
            };

            actual = db.Search(requests);
            Assert.AreEqual("dn: cn=taro.yamada,ou=Employee,dc=example,dc=com", actual.ElementAt(0));
            Assert.AreEqual("objectClass: inetOrgPerson", actual.ElementAt(1));
            Assert.AreEqual("objectClass: posixAccount", actual.ElementAt(2));
            Assert.AreEqual("cn: Taro Yamada", actual.ElementAt(3));
            Assert.AreEqual("uid: taro.yamada", actual.ElementAt(4));
            Assert.AreEqual("uidNumber: 3001", actual.ElementAt(5));
        }

        [Test]
        public void TestCompare()
        {
            var source = new Schema
            {
                { "cn=taro.yamada,ou=Employee,dc=example,dc=com", new Contents {
                    { "dn", new List<string>(){"cn=taro.yamada,ou=Employee,dc=example,dc=com"} },
                    { "objectClass", new List<string>(){ "inetOrgPerson", "posixAccount" } },
                    { "cn", new List<string>(){ "Taro Yamada" } },
                    { "sn", new List<string>(){ "Yamada" } },
                    { "uid", new List<string>(){ "taro.yamada" } },
                    { "userPassword", new List<string>(){ "{SSHA}46mXVpqhvxX3mF+yAzawY47d6ldDwwAs" } },
                    { "uidNumber", new List<string>(){ "3001" } },
                    { "gidNumber", new List<string>(){ "3000" } },
                    { "loginShell", new List<string>(){ "/bin/bash" } },
                    { "homeDirectory", new List<string>(){ "/home/taro.yamada" } },
                } },
            };

            var json = JsonSerializer.Serialize(source);
            stream.Write(Encoding.UTF8.GetBytes(json));

            stream.Position = 0;

            var requests = new Contents()
            {
                { "msgid", new List<string>(){ "1" } },
                { "suffix", new List<string>(){ "dc=example,dc=com" } },
                { "dn", new List<string>(){ "cn=taro.yamada,ou=Employee,dc=example,dc=com" } },
                { "gidNumber", new List<string>(){ "3000" } },
            };

            var actual = db.Compare(requests);
            Assert.AreEqual("RESULT", actual.ElementAt(0));
            Assert.AreEqual("code: 6", actual.ElementAt(1));
        }
    }
}
