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

    // Complete the plusMinus function below.
    static void plusMinus(int[] arr)
    {
        int positives = 0;
        int negatives = 0;
        int zeros = 0;

        foreach (int number in arr)
        {
            if (number > 0)
            {
                positives++;
            }
            else if (number == 0)
            {
                zeros++;
            }
            else
            {
                negatives++;
            }
        }
        decimal occurencePossitive = positives != 0 ? (decimal)positives / arr.Length : 0;
        decimal occurenceNegative = negatives != 0 ? (decimal)negatives / arr.Length : 0;
        decimal occurenceZeros = zeros != 0 ? (decimal)zeros / arr.Length : 0;
        Console.WriteLine(Decimal.Round(occurencePossitive, 6));
        Console.WriteLine(Decimal.Round(occurenceNegative, 6));
        Console.WriteLine(Decimal.Round(occurenceZeros, 6));


    }

    static void Main(string[] args)
    {
        int n = Convert.ToInt32(Console.ReadLine());

        int[] arr = Array.ConvertAll(Console.ReadLine().Split(' '), arrTemp => Convert.ToInt32(arrTemp))
        ;
        plusMinus(arr);
    }
}
