using System.CodeDom.Compiler;
using System.Collections.Generic;
using System.Collections;
using System.ComponentModel;
using System.Diagnostics.CodeAnalysis;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Reflection;
using System.Runtime.Serialization;
using System.Text.RegularExpressions;
using System.Text;
using System;

class Solution
{

    // Complete the compareTriplets function below.
    static List<int> compareTriplets(List<int> a, List<int> b)
    {
        int aResult = 0;
        int bResult = 0;

        if (a.Count > b.Count)
        {
            for (int i = 0; i < a.Count; i++)
            {
                int aVal = a[i];
                int bVal = 0;
                if (i < b.Count)
                {
                    bVal = b[i];
                }
                if (aVal > bVal)
                {
                    aResult++;
                }
                else if (bVal > aVal)
                {
                    bResult++;
                }
            }
        }
        else
        {
            for (int i = 0; i < b.Count; i++)
            {
                int bVal = b[i];
                int aVal = 0;
                if (i < a.Count)
                {
                    aVal = a[i];
                }
                if (aVal > bVal)
                {
                    aResult++;
                }
                else if (bVal > aVal)
                {
                    bResult++;
                }
            }
        }

        return new List<int>() { aResult, bResult };
    }

    static void Main(string[] args)
    {
        TextWriter textWriter = new StreamWriter(@System.Environment.GetEnvironmentVariable("OUTPUT_PATH"), true);

        List<int> a = Console.ReadLine().TrimEnd().Split(' ').ToList().Select(aTemp => Convert.ToInt32(aTemp)).ToList();

        List<int> b = Console.ReadLine().TrimEnd().Split(' ').ToList().Select(bTemp => Convert.ToInt32(bTemp)).ToList();

        List<int> result = compareTriplets(a, b);

        textWriter.WriteLine(String.Join(" ", result));

        textWriter.Flush();
        textWriter.Close();
    }
}
