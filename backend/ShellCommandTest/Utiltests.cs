using NUnit.Framework;
using ShellCommand;
using System.IO;
using System.Linq;
using System.Text;

namespace ShellCommandTest
{
    public class UtilTests
    {
        [SetUp]
        public void Setup()
        {
        }

        [Test]
        public void TestReadInput()
        {
            var input = @"ADD
dn: dc=example,dc=com
msgid: 1
suffix: dc=example,dc=com
objectClass: inetOrgPerson
objectClass: posixAccount
cn: Taro Yamada
sn: Yamada
uid: taro.yamada
userPassword: {SSHA}46mXVpqhvxX3mF+yAzawY47d6ldDwwAs
uidNumber: 3001
gidNumber: 3000
loginShell: /bin/bash
homeDirectory: /home/taro.yamada";

            var stream = new MemoryStream(Encoding.UTF8.GetBytes(input));
            (var command, var request) = Util.ReadInput(stream);

            Assert.AreEqual("ADD", command);
            Assert.AreEqual(1, request["dn"].Count());
            Assert.AreEqual("dc=example,dc=com", request["dn"].ElementAt(0));
            Assert.AreEqual(1, request["msgid"].Count());
            Assert.AreEqual("1", request["msgid"].ElementAt(0));
            Assert.AreEqual(1, request["suffix"].Count());
            Assert.AreEqual("dc=example,dc=com", request["suffix"].ElementAt(0));
            Assert.AreEqual(2, request["objectClass"].Count());
            Assert.AreEqual("inetOrgPerson", request["objectClass"].ElementAt(0));
            Assert.AreEqual("posixAccount", request["objectClass"].ElementAt(1));
            Assert.AreEqual(1, request["cn"].Count());
            Assert.AreEqual("Taro Yamada", request["cn"].ElementAt(0));
            Assert.AreEqual(1, request["sn"].Count());
            Assert.AreEqual("Yamada", request["sn"].ElementAt(0));
            Assert.AreEqual(1, request["uid"].Count());
            Assert.AreEqual("taro.yamada", request["uid"].ElementAt(0));
            Assert.AreEqual(1, request["userPassword"].Count());
            Assert.AreEqual("{SSHA}46mXVpqhvxX3mF+yAzawY47d6ldDwwAs", request["userPassword"].ElementAt(0));
            Assert.AreEqual(1, request["uidNumber"].Count());
            Assert.AreEqual("3001", request["uidNumber"].ElementAt(0));
            Assert.AreEqual(1, request["gidNumber"].Count());
            Assert.AreEqual("3000", request["gidNumber"].ElementAt(0));
            Assert.AreEqual(1, request["loginShell"].Count());
            Assert.AreEqual("/bin/bash", request["loginShell"].ElementAt(0));
            Assert.AreEqual(1, request["homeDirectory"].Count());
            Assert.AreEqual("/home/taro.yamada", request["homeDirectory"].ElementAt(0));
        }

        [Test]
        public void TestWriteOutput()
        {
            var output = new string[]
            {
                "dn: dc=example,dc=com",
                "objectClass: inetOrgPerson",
                "objectClass: posixAccount",
                "cn: Taro Yamada",
                "sn: Yamada",
                "uid: taro.yamada",
                "userPassword: {SSHA}46mXVpqhvxX3mF+yAzawY47d6ldDwwAs",
                "uidNumber: 3001",
                "gidNumber: 3000",
                "loginShell: /bin/bash",
                "homeDirectory: /home/taro.yamada",
            };

            var stream = new MemoryStream();
            Util.WriteOutput(stream, output, true);
            stream.Position = 0;

            using (var reader = new StreamReader(stream))
            {
                Assert.AreEqual("dn: dc=example,dc=com", reader.ReadLine());
                Assert.AreEqual("objectClass: inetOrgPerson", reader.ReadLine());
                Assert.AreEqual("objectClass: posixAccount", reader.ReadLine());
                Assert.AreEqual("cn: Taro Yamada", reader.ReadLine());
                Assert.AreEqual("sn: Yamada", reader.ReadLine());
                Assert.AreEqual("uid: taro.yamada", reader.ReadLine());
                Assert.AreEqual("userPassword: {SSHA}46mXVpqhvxX3mF+yAzawY47d6ldDwwAs", reader.ReadLine());
                Assert.AreEqual("uidNumber: 3001", reader.ReadLine());
                Assert.AreEqual("gidNumber: 3000", reader.ReadLine());
                Assert.AreEqual("loginShell: /bin/bash", reader.ReadLine());
                Assert.AreEqual("homeDirectory: /home/taro.yamada", reader.ReadLine());
            }

        }
    }
}