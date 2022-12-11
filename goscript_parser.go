// Code generated from GoScript.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // GoScript

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 56, 363,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 65, 10, 2, 12, 2,
	14, 2, 68, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4,
	3, 4, 7, 4, 80, 10, 4, 12, 4, 14, 4, 83, 11, 4, 3, 4, 3, 4, 3, 5, 3, 5,
	3, 5, 3, 6, 3, 6, 7, 6, 92, 10, 6, 12, 6, 14, 6, 95, 11, 6, 3, 6, 3, 6,
	3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 103, 10, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9,
	3, 9, 7, 9, 111, 10, 9, 12, 9, 14, 9, 114, 11, 9, 3, 10, 3, 10, 3, 10,
	5, 10, 119, 10, 10, 3, 11, 3, 11, 3, 11, 5, 11, 124, 10, 11, 3, 12, 3,
	12, 3, 12, 3, 12, 7, 12, 130, 10, 12, 12, 12, 14, 12, 133, 11, 12, 5, 12,
	135, 10, 12, 3, 12, 3, 12, 3, 13, 3, 13, 5, 13, 141, 10, 13, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 7, 14, 148, 10, 14, 12, 14, 14, 14, 151, 11, 14,
	3, 14, 3, 14, 3, 15, 3, 15, 5, 15, 157, 10, 15, 3, 15, 3, 15, 5, 15, 161,
	10, 15, 3, 15, 3, 15, 5, 15, 165, 10, 15, 3, 15, 3, 15, 5, 15, 169, 10,
	15, 5, 15, 171, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 177, 10, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3,
	18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 198,
	10, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20,
	208, 10, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5,
	20, 218, 10, 20, 3, 21, 5, 21, 221, 10, 21, 3, 21, 3, 21, 5, 21, 225, 10,
	21, 3, 21, 3, 21, 5, 21, 229, 10, 21, 3, 22, 3, 22, 5, 22, 233, 10, 22,
	3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 5, 24, 248, 10, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 5, 24, 266, 10, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24,
	295, 10, 24, 3, 24, 7, 24, 298, 10, 24, 12, 24, 14, 24, 301, 11, 24, 3,
	25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 5, 25, 309, 10, 25, 3, 26, 3, 26,
	3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 317, 10, 26, 3, 27, 3, 27, 3, 28, 3,
	28, 3, 28, 7, 28, 324, 10, 28, 12, 28, 14, 28, 327, 11, 28, 3, 29, 3, 29,
	3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3,
	29, 3, 29, 3, 29, 5, 29, 344, 10, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29,
	5, 29, 351, 10, 29, 3, 29, 3, 29, 5, 29, 355, 10, 29, 3, 30, 3, 30, 3,
	30, 3, 30, 5, 30, 361, 10, 30, 3, 30, 2, 3, 46, 31, 2, 4, 6, 8, 10, 12,
	14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48,
	50, 52, 54, 56, 58, 2, 7, 3, 2, 14, 18, 3, 2, 41, 42, 3, 2, 43, 45, 3,
	2, 49, 50, 3, 2, 37, 39, 2, 402, 2, 66, 3, 2, 2, 2, 4, 71, 3, 2, 2, 2,
	6, 77, 3, 2, 2, 2, 8, 86, 3, 2, 2, 2, 10, 89, 3, 2, 2, 2, 12, 102, 3, 2,
	2, 2, 14, 104, 3, 2, 2, 2, 16, 107, 3, 2, 2, 2, 18, 115, 3, 2, 2, 2, 20,
	123, 3, 2, 2, 2, 22, 125, 3, 2, 2, 2, 24, 140, 3, 2, 2, 2, 26, 142, 3,
	2, 2, 2, 28, 170, 3, 2, 2, 2, 30, 172, 3, 2, 2, 2, 32, 182, 3, 2, 2, 2,
	34, 187, 3, 2, 2, 2, 36, 189, 3, 2, 2, 2, 38, 217, 3, 2, 2, 2, 40, 220,
	3, 2, 2, 2, 42, 232, 3, 2, 2, 2, 44, 234, 3, 2, 2, 2, 46, 247, 3, 2, 2,
	2, 48, 308, 3, 2, 2, 2, 50, 316, 3, 2, 2, 2, 52, 318, 3, 2, 2, 2, 54, 320,
	3, 2, 2, 2, 56, 354, 3, 2, 2, 2, 58, 360, 3, 2, 2, 2, 60, 65, 5, 4, 3,
	2, 61, 62, 5, 14, 8, 2, 62, 63, 7, 3, 2, 2, 63, 65, 3, 2, 2, 2, 64, 60,
	3, 2, 2, 2, 64, 61, 3, 2, 2, 2, 65, 68, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2,
	66, 67, 3, 2, 2, 2, 67, 69, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 69, 70, 7,
	2, 2, 3, 70, 3, 3, 2, 2, 2, 71, 72, 7, 4, 2, 2, 72, 73, 7, 34, 2, 2, 73,
	74, 5, 6, 4, 2, 74, 75, 5, 28, 15, 2, 75, 76, 5, 10, 6, 2, 76, 5, 3, 2,
	2, 2, 77, 81, 7, 5, 2, 2, 78, 80, 5, 8, 5, 2, 79, 78, 3, 2, 2, 2, 80, 83,
	3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 84, 3, 2, 2, 2,
	83, 81, 3, 2, 2, 2, 84, 85, 7, 6, 2, 2, 85, 7, 3, 2, 2, 2, 86, 87, 5, 28,
	15, 2, 87, 88, 7, 34, 2, 2, 88, 9, 3, 2, 2, 2, 89, 93, 7, 7, 2, 2, 90,
	92, 5, 12, 7, 2, 91, 90, 3, 2, 2, 2, 92, 95, 3, 2, 2, 2, 93, 91, 3, 2,
	2, 2, 93, 94, 3, 2, 2, 2, 94, 96, 3, 2, 2, 2, 95, 93, 3, 2, 2, 2, 96, 97,
	7, 8, 2, 2, 97, 11, 3, 2, 2, 2, 98, 99, 5, 14, 8, 2, 99, 100, 7, 3, 2,
	2, 100, 103, 3, 2, 2, 2, 101, 103, 5, 38, 20, 2, 102, 98, 3, 2, 2, 2, 102,
	101, 3, 2, 2, 2, 103, 13, 3, 2, 2, 2, 104, 105, 5, 28, 15, 2, 105, 106,
	5, 16, 9, 2, 106, 15, 3, 2, 2, 2, 107, 112, 5, 18, 10, 2, 108, 109, 7,
	9, 2, 2, 109, 111, 5, 18, 10, 2, 110, 108, 3, 2, 2, 2, 111, 114, 3, 2,
	2, 2, 112, 110, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113, 17, 3, 2, 2, 2,
	114, 112, 3, 2, 2, 2, 115, 118, 7, 34, 2, 2, 116, 117, 7, 48, 2, 2, 117,
	119, 5, 20, 11, 2, 118, 116, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 19,
	3, 2, 2, 2, 120, 124, 5, 22, 12, 2, 121, 124, 5, 26, 14, 2, 122, 124, 5,
	46, 24, 2, 123, 120, 3, 2, 2, 2, 123, 121, 3, 2, 2, 2, 123, 122, 3, 2,
	2, 2, 124, 21, 3, 2, 2, 2, 125, 134, 7, 7, 2, 2, 126, 131, 5, 24, 13, 2,
	127, 128, 7, 9, 2, 2, 128, 130, 5, 24, 13, 2, 129, 127, 3, 2, 2, 2, 130,
	133, 3, 2, 2, 2, 131, 129, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 135,
	3, 2, 2, 2, 133, 131, 3, 2, 2, 2, 134, 126, 3, 2, 2, 2, 134, 135, 3, 2,
	2, 2, 135, 136, 3, 2, 2, 2, 136, 137, 7, 8, 2, 2, 137, 23, 3, 2, 2, 2,
	138, 141, 5, 26, 14, 2, 139, 141, 5, 46, 24, 2, 140, 138, 3, 2, 2, 2, 140,
	139, 3, 2, 2, 2, 141, 25, 3, 2, 2, 2, 142, 149, 7, 7, 2, 2, 143, 144, 5,
	46, 24, 2, 144, 145, 7, 10, 2, 2, 145, 146, 5, 20, 11, 2, 146, 148, 3,
	2, 2, 2, 147, 143, 3, 2, 2, 2, 148, 151, 3, 2, 2, 2, 149, 147, 3, 2, 2,
	2, 149, 150, 3, 2, 2, 2, 150, 152, 3, 2, 2, 2, 151, 149, 3, 2, 2, 2, 152,
	153, 7, 8, 2, 2, 153, 27, 3, 2, 2, 2, 154, 156, 5, 36, 19, 2, 155, 157,
	7, 35, 2, 2, 156, 155, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 171, 3, 2,
	2, 2, 158, 160, 5, 30, 16, 2, 159, 161, 7, 35, 2, 2, 160, 159, 3, 2, 2,
	2, 160, 161, 3, 2, 2, 2, 161, 171, 3, 2, 2, 2, 162, 164, 5, 32, 17, 2,
	163, 165, 7, 35, 2, 2, 164, 163, 3, 2, 2, 2, 164, 165, 3, 2, 2, 2, 165,
	171, 3, 2, 2, 2, 166, 168, 5, 34, 18, 2, 167, 169, 7, 35, 2, 2, 168, 167,
	3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 171, 3, 2, 2, 2, 170, 154, 3, 2,
	2, 2, 170, 158, 3, 2, 2, 2, 170, 162, 3, 2, 2, 2, 170, 166, 3, 2, 2, 2,
	171, 29, 3, 2, 2, 2, 172, 173, 7, 11, 2, 2, 173, 176, 7, 47, 2, 2, 174,
	177, 5, 36, 19, 2, 175, 177, 5, 32, 17, 2, 176, 174, 3, 2, 2, 2, 176, 175,
	3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 179, 7, 9, 2, 2, 179, 180, 5, 28,
	15, 2, 180, 181, 7, 46, 2, 2, 181, 31, 3, 2, 2, 2, 182, 183, 7, 12, 2,
	2, 183, 184, 7, 47, 2, 2, 184, 185, 7, 34, 2, 2, 185, 186, 7, 46, 2, 2,
	186, 33, 3, 2, 2, 2, 187, 188, 7, 13, 2, 2, 188, 35, 3, 2, 2, 2, 189, 190,
	9, 2, 2, 2, 190, 37, 3, 2, 2, 2, 191, 218, 5, 10, 6, 2, 192, 193, 7, 19,
	2, 2, 193, 194, 5, 46, 24, 2, 194, 197, 5, 38, 20, 2, 195, 196, 7, 20,
	2, 2, 196, 198, 5, 38, 20, 2, 197, 195, 3, 2, 2, 2, 197, 198, 3, 2, 2,
	2, 198, 218, 3, 2, 2, 2, 199, 200, 7, 21, 2, 2, 200, 201, 7, 5, 2, 2, 201,
	202, 5, 40, 21, 2, 202, 203, 7, 6, 2, 2, 203, 204, 5, 38, 20, 2, 204, 218,
	3, 2, 2, 2, 205, 207, 7, 22, 2, 2, 206, 208, 5, 46, 24, 2, 207, 206, 3,
	2, 2, 2, 207, 208, 3, 2, 2, 2, 208, 209, 3, 2, 2, 2, 209, 218, 7, 3, 2,
	2, 210, 211, 7, 23, 2, 2, 211, 218, 7, 3, 2, 2, 212, 213, 7, 24, 2, 2,
	213, 218, 7, 3, 2, 2, 214, 215, 5, 46, 24, 2, 215, 216, 7, 3, 2, 2, 216,
	218, 3, 2, 2, 2, 217, 191, 3, 2, 2, 2, 217, 192, 3, 2, 2, 2, 217, 199,
	3, 2, 2, 2, 217, 205, 3, 2, 2, 2, 217, 210, 3, 2, 2, 2, 217, 212, 3, 2,
	2, 2, 217, 214, 3, 2, 2, 2, 218, 39, 3, 2, 2, 2, 219, 221, 5, 42, 22, 2,
	220, 219, 3, 2, 2, 2, 220, 221, 3, 2, 2, 2, 221, 222, 3, 2, 2, 2, 222,
	224, 7, 3, 2, 2, 223, 225, 5, 46, 24, 2, 224, 223, 3, 2, 2, 2, 224, 225,
	3, 2, 2, 2, 225, 226, 3, 2, 2, 2, 226, 228, 7, 3, 2, 2, 227, 229, 5, 44,
	23, 2, 228, 227, 3, 2, 2, 2, 228, 229, 3, 2, 2, 2, 229, 41, 3, 2, 2, 2,
	230, 233, 5, 14, 8, 2, 231, 233, 5, 54, 28, 2, 232, 230, 3, 2, 2, 2, 232,
	231, 3, 2, 2, 2, 233, 43, 3, 2, 2, 2, 234, 235, 5, 54, 28, 2, 235, 45,
	3, 2, 2, 2, 236, 237, 8, 24, 1, 2, 237, 248, 5, 48, 25, 2, 238, 239, 9,
	3, 2, 2, 239, 248, 5, 46, 24, 14, 240, 241, 7, 51, 2, 2, 241, 248, 5, 46,
	24, 13, 242, 243, 7, 28, 2, 2, 243, 248, 5, 56, 29, 2, 244, 245, 7, 34,
	2, 2, 245, 246, 7, 32, 2, 2, 246, 248, 5, 46, 24, 3, 247, 236, 3, 2, 2,
	2, 247, 238, 3, 2, 2, 2, 247, 240, 3, 2, 2, 2, 247, 242, 3, 2, 2, 2, 247,
	244, 3, 2, 2, 2, 248, 299, 3, 2, 2, 2, 249, 250, 12, 17, 2, 2, 250, 251,
	7, 25, 2, 2, 251, 298, 5, 46, 24, 18, 252, 253, 12, 11, 2, 2, 253, 254,
	9, 4, 2, 2, 254, 298, 5, 46, 24, 12, 255, 256, 12, 10, 2, 2, 256, 257,
	9, 3, 2, 2, 257, 298, 5, 46, 24, 11, 258, 265, 12, 9, 2, 2, 259, 260, 7,
	47, 2, 2, 260, 266, 7, 48, 2, 2, 261, 262, 7, 46, 2, 2, 262, 266, 7, 48,
	2, 2, 263, 266, 7, 46, 2, 2, 264, 266, 7, 47, 2, 2, 265, 259, 3, 2, 2,
	2, 265, 261, 3, 2, 2, 2, 265, 263, 3, 2, 2, 2, 265, 264, 3, 2, 2, 2, 266,
	267, 3, 2, 2, 2, 267, 298, 5, 46, 24, 10, 268, 269, 12, 8, 2, 2, 269, 270,
	9, 5, 2, 2, 270, 298, 5, 46, 24, 9, 271, 272, 12, 7, 2, 2, 272, 273, 7,
	29, 2, 2, 273, 298, 5, 46, 24, 8, 274, 275, 12, 6, 2, 2, 275, 276, 7, 30,
	2, 2, 276, 298, 5, 46, 24, 7, 277, 278, 12, 5, 2, 2, 278, 279, 7, 31, 2,
	2, 279, 280, 5, 46, 24, 2, 280, 281, 7, 10, 2, 2, 281, 282, 5, 46, 24,
	6, 282, 298, 3, 2, 2, 2, 283, 284, 12, 4, 2, 2, 284, 285, 7, 48, 2, 2,
	285, 298, 5, 46, 24, 4, 286, 287, 12, 16, 2, 2, 287, 288, 7, 26, 2, 2,
	288, 289, 5, 46, 24, 2, 289, 290, 7, 27, 2, 2, 290, 298, 3, 2, 2, 2, 291,
	292, 12, 15, 2, 2, 292, 294, 7, 5, 2, 2, 293, 295, 5, 54, 28, 2, 294, 293,
	3, 2, 2, 2, 294, 295, 3, 2, 2, 2, 295, 296, 3, 2, 2, 2, 296, 298, 7, 6,
	2, 2, 297, 249, 3, 2, 2, 2, 297, 252, 3, 2, 2, 2, 297, 255, 3, 2, 2, 2,
	297, 258, 3, 2, 2, 2, 297, 268, 3, 2, 2, 2, 297, 271, 3, 2, 2, 2, 297,
	274, 3, 2, 2, 2, 297, 277, 3, 2, 2, 2, 297, 283, 3, 2, 2, 2, 297, 286,
	3, 2, 2, 2, 297, 291, 3, 2, 2, 2, 298, 301, 3, 2, 2, 2, 299, 297, 3, 2,
	2, 2, 299, 300, 3, 2, 2, 2, 300, 47, 3, 2, 2, 2, 301, 299, 3, 2, 2, 2,
	302, 303, 7, 5, 2, 2, 303, 304, 5, 46, 24, 2, 304, 305, 7, 6, 2, 2, 305,
	309, 3, 2, 2, 2, 306, 309, 5, 50, 26, 2, 307, 309, 7, 34, 2, 2, 308, 302,
	3, 2, 2, 2, 308, 306, 3, 2, 2, 2, 308, 307, 3, 2, 2, 2, 309, 49, 3, 2,
	2, 2, 310, 317, 5, 52, 27, 2, 311, 317, 7, 40, 2, 2, 312, 317, 7, 52, 2,
	2, 313, 317, 7, 53, 2, 2, 314, 317, 7, 33, 2, 2, 315, 317, 7, 36, 2, 2,
	316, 310, 3, 2, 2, 2, 316, 311, 3, 2, 2, 2, 316, 312, 3, 2, 2, 2, 316,
	313, 3, 2, 2, 2, 316, 314, 3, 2, 2, 2, 316, 315, 3, 2, 2, 2, 317, 51, 3,
	2, 2, 2, 318, 319, 9, 6, 2, 2, 319, 53, 3, 2, 2, 2, 320, 325, 5, 46, 24,
	2, 321, 322, 7, 9, 2, 2, 322, 324, 5, 46, 24, 2, 323, 321, 3, 2, 2, 2,
	324, 327, 3, 2, 2, 2, 325, 323, 3, 2, 2, 2, 325, 326, 3, 2, 2, 2, 326,
	55, 3, 2, 2, 2, 327, 325, 3, 2, 2, 2, 328, 329, 5, 30, 16, 2, 329, 330,
	5, 26, 14, 2, 330, 355, 3, 2, 2, 2, 331, 332, 5, 58, 30, 2, 332, 333, 7,
	26, 2, 2, 333, 334, 7, 27, 2, 2, 334, 335, 5, 22, 12, 2, 335, 355, 3, 2,
	2, 2, 336, 337, 5, 32, 17, 2, 337, 338, 7, 5, 2, 2, 338, 339, 7, 6, 2,
	2, 339, 355, 3, 2, 2, 2, 340, 341, 5, 36, 19, 2, 341, 343, 7, 5, 2, 2,
	342, 344, 5, 46, 24, 2, 343, 342, 3, 2, 2, 2, 343, 344, 3, 2, 2, 2, 344,
	345, 3, 2, 2, 2, 345, 346, 7, 6, 2, 2, 346, 355, 3, 2, 2, 2, 347, 348,
	5, 34, 18, 2, 348, 350, 7, 5, 2, 2, 349, 351, 5, 46, 24, 2, 350, 349, 3,
	2, 2, 2, 350, 351, 3, 2, 2, 2, 351, 352, 3, 2, 2, 2, 352, 353, 7, 6, 2,
	2, 353, 355, 3, 2, 2, 2, 354, 328, 3, 2, 2, 2, 354, 331, 3, 2, 2, 2, 354,
	336, 3, 2, 2, 2, 354, 340, 3, 2, 2, 2, 354, 347, 3, 2, 2, 2, 355, 57, 3,
	2, 2, 2, 356, 361, 5, 36, 19, 2, 357, 361, 5, 30, 16, 2, 358, 361, 5, 32,
	17, 2, 359, 361, 5, 34, 18, 2, 360, 356, 3, 2, 2, 2, 360, 357, 3, 2, 2,
	2, 360, 358, 3, 2, 2, 2, 360, 359, 3, 2, 2, 2, 361, 59, 3, 2, 2, 2, 39,
	64, 66, 81, 93, 102, 112, 118, 123, 131, 134, 140, 149, 156, 160, 164,
	168, 170, 176, 197, 207, 217, 220, 224, 228, 232, 247, 265, 294, 297, 299,
	308, 316, 325, 343, 350, 354, 360,
}
var literalNames = []string{
	"", "';'", "'func'", "'('", "')'", "'{'", "'}'", "','", "':'", "'map'",
	"'connector'", "'dynamic'", "'int'", "'float'", "'bool'", "'char'", "'string'",
	"'if'", "'else'", "'for'", "'return'", "'break'", "'continue'", "'.'",
	"'['", "']'", "'new'", "'&&'", "'||'", "'?'", "':='", "", "", "", "'null'",
	"", "", "", "", "'+'", "'-'", "'*'", "'/'", "'%'", "'>'", "'<'", "'='",
	"'=='", "'!='", "'!'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "BooleanLiteral", "Identifier",
	"Brackets", "NULL", "HexLiteral", "DecimalLiteral", "OctalLiteral", "FloatingPointLiteral",
	"Plus", "Minus", "Multiplication", "Division", "Percent", "Gt", "Lt", "Assign",
	"Eq", "NEq", "Not", "CharacterLiteral", "StringLiteral", "COMMENT", "WS",
	"LINE_COMMENT",
}

var ruleNames = []string{
	"compilationUnit", "functionDeclaration", "formalParameters", "formalParameterDecl",
	"block", "blockStatement", "variableDeclaration", "variableDeclarators",
	"variableDeclarator", "variableInitializer", "arrayInitializer", "arrayInitializerElement",
	"mapInitializer", "type_", "mapType", "connectorType", "dynamicType", "primitiveType",
	"statement", "forControl", "forInit", "forUpdate", "expression", "primary",
	"literal", "integerLiteral", "expressionList", "creator", "creatorName",
}

type GoScriptParser struct {
	*antlr.BaseParser
}

// NewGoScriptParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *GoScriptParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewGoScriptParser(input antlr.TokenStream) *GoScriptParser {
	this := new(GoScriptParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "GoScript.g4"

	return this
}

// GoScriptParser tokens.
const (
	GoScriptParserEOF                  = antlr.TokenEOF
	GoScriptParserT__0                 = 1
	GoScriptParserT__1                 = 2
	GoScriptParserT__2                 = 3
	GoScriptParserT__3                 = 4
	GoScriptParserT__4                 = 5
	GoScriptParserT__5                 = 6
	GoScriptParserT__6                 = 7
	GoScriptParserT__7                 = 8
	GoScriptParserT__8                 = 9
	GoScriptParserT__9                 = 10
	GoScriptParserT__10                = 11
	GoScriptParserT__11                = 12
	GoScriptParserT__12                = 13
	GoScriptParserT__13                = 14
	GoScriptParserT__14                = 15
	GoScriptParserT__15                = 16
	GoScriptParserT__16                = 17
	GoScriptParserT__17                = 18
	GoScriptParserT__18                = 19
	GoScriptParserT__19                = 20
	GoScriptParserT__20                = 21
	GoScriptParserT__21                = 22
	GoScriptParserT__22                = 23
	GoScriptParserT__23                = 24
	GoScriptParserT__24                = 25
	GoScriptParserT__25                = 26
	GoScriptParserT__26                = 27
	GoScriptParserT__27                = 28
	GoScriptParserT__28                = 29
	GoScriptParserT__29                = 30
	GoScriptParserBooleanLiteral       = 31
	GoScriptParserIdentifier           = 32
	GoScriptParserBrackets             = 33
	GoScriptParserNULL                 = 34
	GoScriptParserHexLiteral           = 35
	GoScriptParserDecimalLiteral       = 36
	GoScriptParserOctalLiteral         = 37
	GoScriptParserFloatingPointLiteral = 38
	GoScriptParserPlus                 = 39
	GoScriptParserMinus                = 40
	GoScriptParserMultiplication       = 41
	GoScriptParserDivision             = 42
	GoScriptParserPercent              = 43
	GoScriptParserGt                   = 44
	GoScriptParserLt                   = 45
	GoScriptParserAssign               = 46
	GoScriptParserEq                   = 47
	GoScriptParserNEq                  = 48
	GoScriptParserNot                  = 49
	GoScriptParserCharacterLiteral     = 50
	GoScriptParserStringLiteral        = 51
	GoScriptParserCOMMENT              = 52
	GoScriptParserWS                   = 53
	GoScriptParserLINE_COMMENT         = 54
)

// GoScriptParser rules.
const (
	GoScriptParserRULE_compilationUnit         = 0
	GoScriptParserRULE_functionDeclaration     = 1
	GoScriptParserRULE_formalParameters        = 2
	GoScriptParserRULE_formalParameterDecl     = 3
	GoScriptParserRULE_block                   = 4
	GoScriptParserRULE_blockStatement          = 5
	GoScriptParserRULE_variableDeclaration     = 6
	GoScriptParserRULE_variableDeclarators     = 7
	GoScriptParserRULE_variableDeclarator      = 8
	GoScriptParserRULE_variableInitializer     = 9
	GoScriptParserRULE_arrayInitializer        = 10
	GoScriptParserRULE_arrayInitializerElement = 11
	GoScriptParserRULE_mapInitializer          = 12
	GoScriptParserRULE_type_                   = 13
	GoScriptParserRULE_mapType                 = 14
	GoScriptParserRULE_connectorType           = 15
	GoScriptParserRULE_dynamicType             = 16
	GoScriptParserRULE_primitiveType           = 17
	GoScriptParserRULE_statement               = 18
	GoScriptParserRULE_forControl              = 19
	GoScriptParserRULE_forInit                 = 20
	GoScriptParserRULE_forUpdate               = 21
	GoScriptParserRULE_expression              = 22
	GoScriptParserRULE_primary                 = 23
	GoScriptParserRULE_literal                 = 24
	GoScriptParserRULE_integerLiteral          = 25
	GoScriptParserRULE_expressionList          = 26
	GoScriptParserRULE_creator                 = 27
	GoScriptParserRULE_creatorName             = 28
)

// ICompilationUnitContext is an interface to support dynamic dispatch.
type ICompilationUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompilationUnitContext differentiates from other interfaces.
	IsCompilationUnitContext()
}

type CompilationUnitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompilationUnitContext() *CompilationUnitContext {
	var p = new(CompilationUnitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_compilationUnit
	return p
}

func (*CompilationUnitContext) IsCompilationUnitContext() {}

func NewCompilationUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompilationUnitContext {
	var p = new(CompilationUnitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_compilationUnit

	return p
}

func (s *CompilationUnitContext) GetParser() antlr.Parser { return s.parser }

func (s *CompilationUnitContext) EOF() antlr.TerminalNode {
	return s.GetToken(GoScriptParserEOF, 0)
}

func (s *CompilationUnitContext) AllFunctionDeclaration() []IFunctionDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem())
	var tst = make([]IFunctionDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionDeclarationContext)
		}
	}

	return tst
}

func (s *CompilationUnitContext) FunctionDeclaration(i int) IFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *CompilationUnitContext) AllVariableDeclaration() []IVariableDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem())
	var tst = make([]IVariableDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableDeclarationContext)
		}
	}

	return tst
}

func (s *CompilationUnitContext) VariableDeclaration(i int) IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *CompilationUnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompilationUnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompilationUnitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitCompilationUnit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) CompilationUnit() (localctx ICompilationUnitContext) {
	localctx = NewCompilationUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GoScriptParserRULE_compilationUnit)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GoScriptParserT__1)|(1<<GoScriptParserT__8)|(1<<GoScriptParserT__9)|(1<<GoScriptParserT__10)|(1<<GoScriptParserT__11)|(1<<GoScriptParserT__12)|(1<<GoScriptParserT__13)|(1<<GoScriptParserT__14)|(1<<GoScriptParserT__15))) != 0 {
		p.SetState(62)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case GoScriptParserT__1:
			{
				p.SetState(58)
				p.FunctionDeclaration()
			}

		case GoScriptParserT__8, GoScriptParserT__9, GoScriptParserT__10, GoScriptParserT__11, GoScriptParserT__12, GoScriptParserT__13, GoScriptParserT__14, GoScriptParserT__15:
			{
				p.SetState(59)
				p.VariableDeclaration()
			}
			{
				p.SetState(60)
				p.Match(GoScriptParserT__0)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(67)
		p.Match(GoScriptParserEOF)
	}

	return localctx
}

// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_functionDeclaration
	return p
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_functionDeclaration

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(GoScriptParserIdentifier, 0)
}

func (s *FunctionDeclarationContext) FormalParameters() IFormalParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormalParametersContext)
}

func (s *FunctionDeclarationContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *FunctionDeclarationContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitFunctionDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) FunctionDeclaration() (localctx IFunctionDeclarationContext) {
	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GoScriptParserRULE_functionDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		p.Match(GoScriptParserT__1)
	}
	{
		p.SetState(70)
		p.Match(GoScriptParserIdentifier)
	}
	{
		p.SetState(71)
		p.FormalParameters()
	}
	{
		p.SetState(72)
		p.Type_()
	}
	{
		p.SetState(73)
		p.Block()
	}

	return localctx
}

// IFormalParametersContext is an interface to support dynamic dispatch.
type IFormalParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormalParametersContext differentiates from other interfaces.
	IsFormalParametersContext()
}

type FormalParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalParametersContext() *FormalParametersContext {
	var p = new(FormalParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_formalParameters
	return p
}

func (*FormalParametersContext) IsFormalParametersContext() {}

func NewFormalParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalParametersContext {
	var p = new(FormalParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_formalParameters

	return p
}

func (s *FormalParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalParametersContext) AllFormalParameterDecl() []IFormalParameterDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFormalParameterDeclContext)(nil)).Elem())
	var tst = make([]IFormalParameterDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFormalParameterDeclContext)
		}
	}

	return tst
}

func (s *FormalParametersContext) FormalParameterDecl(i int) IFormalParameterDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalParameterDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFormalParameterDeclContext)
}

func (s *FormalParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormalParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitFormalParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) FormalParameters() (localctx IFormalParametersContext) {
	localctx = NewFormalParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GoScriptParserRULE_formalParameters)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		p.Match(GoScriptParserT__2)
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GoScriptParserT__8)|(1<<GoScriptParserT__9)|(1<<GoScriptParserT__10)|(1<<GoScriptParserT__11)|(1<<GoScriptParserT__12)|(1<<GoScriptParserT__13)|(1<<GoScriptParserT__14)|(1<<GoScriptParserT__15))) != 0 {
		{
			p.SetState(76)
			p.FormalParameterDecl()
		}

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(82)
		p.Match(GoScriptParserT__3)
	}

	return localctx
}

// IFormalParameterDeclContext is an interface to support dynamic dispatch.
type IFormalParameterDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormalParameterDeclContext differentiates from other interfaces.
	IsFormalParameterDeclContext()
}

type FormalParameterDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalParameterDeclContext() *FormalParameterDeclContext {
	var p = new(FormalParameterDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_formalParameterDecl
	return p
}

func (*FormalParameterDeclContext) IsFormalParameterDeclContext() {}

func NewFormalParameterDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalParameterDeclContext {
	var p = new(FormalParameterDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_formalParameterDecl

	return p
}

func (s *FormalParameterDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalParameterDeclContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *FormalParameterDeclContext) Identifier() antlr.TerminalNode {
	return s.GetToken(GoScriptParserIdentifier, 0)
}

func (s *FormalParameterDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalParameterDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormalParameterDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitFormalParameterDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) FormalParameterDecl() (localctx IFormalParameterDeclContext) {
	localctx = NewFormalParameterDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GoScriptParserRULE_formalParameterDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.Type_()
	}
	{
		p.SetState(85)
		p.Match(GoScriptParserIdentifier)
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllBlockStatement() []IBlockStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockStatementContext)(nil)).Elem())
	var tst = make([]IBlockStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockStatementContext)
		}
	}

	return tst
}

func (s *BlockContext) BlockStatement(i int) IBlockStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockStatementContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GoScriptParserRULE_block)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(87)
		p.Match(GoScriptParserT__4)
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GoScriptParserT__2)|(1<<GoScriptParserT__4)|(1<<GoScriptParserT__8)|(1<<GoScriptParserT__9)|(1<<GoScriptParserT__10)|(1<<GoScriptParserT__11)|(1<<GoScriptParserT__12)|(1<<GoScriptParserT__13)|(1<<GoScriptParserT__14)|(1<<GoScriptParserT__15)|(1<<GoScriptParserT__16)|(1<<GoScriptParserT__18)|(1<<GoScriptParserT__19)|(1<<GoScriptParserT__20)|(1<<GoScriptParserT__21)|(1<<GoScriptParserT__25)|(1<<GoScriptParserBooleanLiteral))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(GoScriptParserIdentifier-32))|(1<<(GoScriptParserNULL-32))|(1<<(GoScriptParserHexLiteral-32))|(1<<(GoScriptParserDecimalLiteral-32))|(1<<(GoScriptParserOctalLiteral-32))|(1<<(GoScriptParserFloatingPointLiteral-32))|(1<<(GoScriptParserPlus-32))|(1<<(GoScriptParserMinus-32))|(1<<(GoScriptParserNot-32))|(1<<(GoScriptParserCharacterLiteral-32))|(1<<(GoScriptParserStringLiteral-32)))) != 0) {
		{
			p.SetState(88)
			p.BlockStatement()
		}

		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(94)
		p.Match(GoScriptParserT__5)
	}

	return localctx
}

// IBlockStatementContext is an interface to support dynamic dispatch.
type IBlockStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockStatementContext differentiates from other interfaces.
	IsBlockStatementContext()
}

type BlockStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockStatementContext() *BlockStatementContext {
	var p = new(BlockStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_blockStatement
	return p
}

func (*BlockStatementContext) IsBlockStatementContext() {}

func NewBlockStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockStatementContext {
	var p = new(BlockStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_blockStatement

	return p
}

func (s *BlockStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockStatementContext) VariableDeclaration() IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *BlockStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitBlockStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) BlockStatement() (localctx IBlockStatementContext) {
	localctx = NewBlockStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GoScriptParserRULE_blockStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(100)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoScriptParserT__8, GoScriptParserT__9, GoScriptParserT__10, GoScriptParserT__11, GoScriptParserT__12, GoScriptParserT__13, GoScriptParserT__14, GoScriptParserT__15:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(96)
			p.VariableDeclaration()
		}
		{
			p.SetState(97)
			p.Match(GoScriptParserT__0)
		}

	case GoScriptParserT__2, GoScriptParserT__4, GoScriptParserT__16, GoScriptParserT__18, GoScriptParserT__19, GoScriptParserT__20, GoScriptParserT__21, GoScriptParserT__25, GoScriptParserBooleanLiteral, GoScriptParserIdentifier, GoScriptParserNULL, GoScriptParserHexLiteral, GoScriptParserDecimalLiteral, GoScriptParserOctalLiteral, GoScriptParserFloatingPointLiteral, GoScriptParserPlus, GoScriptParserMinus, GoScriptParserNot, GoScriptParserCharacterLiteral, GoScriptParserStringLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(99)
			p.Statement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVariableDeclarationContext is an interface to support dynamic dispatch.
type IVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclarationContext differentiates from other interfaces.
	IsVariableDeclarationContext()
}

type VariableDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationContext() *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_variableDeclaration
	return p
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *VariableDeclarationContext) VariableDeclarators() IVariableDeclaratorsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclaratorsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorsContext)
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitVariableDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GoScriptParserRULE_variableDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		p.Type_()
	}
	{
		p.SetState(103)
		p.VariableDeclarators()
	}

	return localctx
}

// IVariableDeclaratorsContext is an interface to support dynamic dispatch.
type IVariableDeclaratorsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclaratorsContext differentiates from other interfaces.
	IsVariableDeclaratorsContext()
}

type VariableDeclaratorsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclaratorsContext() *VariableDeclaratorsContext {
	var p = new(VariableDeclaratorsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_variableDeclarators
	return p
}

func (*VariableDeclaratorsContext) IsVariableDeclaratorsContext() {}

func NewVariableDeclaratorsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclaratorsContext {
	var p = new(VariableDeclaratorsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_variableDeclarators

	return p
}

func (s *VariableDeclaratorsContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclaratorsContext) AllVariableDeclarator() []IVariableDeclaratorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableDeclaratorContext)(nil)).Elem())
	var tst = make([]IVariableDeclaratorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableDeclaratorContext)
		}
	}

	return tst
}

func (s *VariableDeclaratorsContext) VariableDeclarator(i int) IVariableDeclaratorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclaratorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclaratorContext)
}

func (s *VariableDeclaratorsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclaratorsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclaratorsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitVariableDeclarators(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) VariableDeclarators() (localctx IVariableDeclaratorsContext) {
	localctx = NewVariableDeclaratorsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GoScriptParserRULE_variableDeclarators)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.VariableDeclarator()
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoScriptParserT__6 {
		{
			p.SetState(106)
			p.Match(GoScriptParserT__6)
		}
		{
			p.SetState(107)
			p.VariableDeclarator()
		}

		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IVariableDeclaratorContext is an interface to support dynamic dispatch.
type IVariableDeclaratorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclaratorContext differentiates from other interfaces.
	IsVariableDeclaratorContext()
}

type VariableDeclaratorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclaratorContext() *VariableDeclaratorContext {
	var p = new(VariableDeclaratorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_variableDeclarator
	return p
}

func (*VariableDeclaratorContext) IsVariableDeclaratorContext() {}

func NewVariableDeclaratorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclaratorContext {
	var p = new(VariableDeclaratorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_variableDeclarator

	return p
}

func (s *VariableDeclaratorContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclaratorContext) Identifier() antlr.TerminalNode {
	return s.GetToken(GoScriptParserIdentifier, 0)
}

func (s *VariableDeclaratorContext) Assign() antlr.TerminalNode {
	return s.GetToken(GoScriptParserAssign, 0)
}

func (s *VariableDeclaratorContext) VariableInitializer() IVariableInitializerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableInitializerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableInitializerContext)
}

func (s *VariableDeclaratorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclaratorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclaratorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitVariableDeclarator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) VariableDeclarator() (localctx IVariableDeclaratorContext) {
	localctx = NewVariableDeclaratorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GoScriptParserRULE_variableDeclarator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)
		p.Match(GoScriptParserIdentifier)
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GoScriptParserAssign {
		{
			p.SetState(114)
			p.Match(GoScriptParserAssign)
		}
		{
			p.SetState(115)
			p.VariableInitializer()
		}

	}

	return localctx
}

// IVariableInitializerContext is an interface to support dynamic dispatch.
type IVariableInitializerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableInitializerContext differentiates from other interfaces.
	IsVariableInitializerContext()
}

type VariableInitializerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableInitializerContext() *VariableInitializerContext {
	var p = new(VariableInitializerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_variableInitializer
	return p
}

func (*VariableInitializerContext) IsVariableInitializerContext() {}

func NewVariableInitializerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableInitializerContext {
	var p = new(VariableInitializerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_variableInitializer

	return p
}

func (s *VariableInitializerContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableInitializerContext) ArrayInitializer() IArrayInitializerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayInitializerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayInitializerContext)
}

func (s *VariableInitializerContext) MapInitializer() IMapInitializerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapInitializerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapInitializerContext)
}

func (s *VariableInitializerContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VariableInitializerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableInitializerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableInitializerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitVariableInitializer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) VariableInitializer() (localctx IVariableInitializerContext) {
	localctx = NewVariableInitializerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GoScriptParserRULE_variableInitializer)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(118)
			p.ArrayInitializer()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(119)
			p.MapInitializer()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(120)
			p.expression(0)
		}

	}

	return localctx
}

// IArrayInitializerContext is an interface to support dynamic dispatch.
type IArrayInitializerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayInitializerContext differentiates from other interfaces.
	IsArrayInitializerContext()
}

type ArrayInitializerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayInitializerContext() *ArrayInitializerContext {
	var p = new(ArrayInitializerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_arrayInitializer
	return p
}

func (*ArrayInitializerContext) IsArrayInitializerContext() {}

func NewArrayInitializerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayInitializerContext {
	var p = new(ArrayInitializerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_arrayInitializer

	return p
}

func (s *ArrayInitializerContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayInitializerContext) AllArrayInitializerElement() []IArrayInitializerElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArrayInitializerElementContext)(nil)).Elem())
	var tst = make([]IArrayInitializerElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArrayInitializerElementContext)
		}
	}

	return tst
}

func (s *ArrayInitializerContext) ArrayInitializerElement(i int) IArrayInitializerElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayInitializerElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArrayInitializerElementContext)
}

func (s *ArrayInitializerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayInitializerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayInitializerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitArrayInitializer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) ArrayInitializer() (localctx IArrayInitializerContext) {
	localctx = NewArrayInitializerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GoScriptParserRULE_arrayInitializer)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.Match(GoScriptParserT__4)
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GoScriptParserT__2)|(1<<GoScriptParserT__4)|(1<<GoScriptParserT__25)|(1<<GoScriptParserBooleanLiteral))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(GoScriptParserIdentifier-32))|(1<<(GoScriptParserNULL-32))|(1<<(GoScriptParserHexLiteral-32))|(1<<(GoScriptParserDecimalLiteral-32))|(1<<(GoScriptParserOctalLiteral-32))|(1<<(GoScriptParserFloatingPointLiteral-32))|(1<<(GoScriptParserPlus-32))|(1<<(GoScriptParserMinus-32))|(1<<(GoScriptParserNot-32))|(1<<(GoScriptParserCharacterLiteral-32))|(1<<(GoScriptParserStringLiteral-32)))) != 0) {
		{
			p.SetState(124)
			p.ArrayInitializerElement()
		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GoScriptParserT__6 {
			{
				p.SetState(125)
				p.Match(GoScriptParserT__6)
			}
			{
				p.SetState(126)
				p.ArrayInitializerElement()
			}

			p.SetState(131)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(134)
		p.Match(GoScriptParserT__5)
	}

	return localctx
}

// IArrayInitializerElementContext is an interface to support dynamic dispatch.
type IArrayInitializerElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayInitializerElementContext differentiates from other interfaces.
	IsArrayInitializerElementContext()
}

type ArrayInitializerElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayInitializerElementContext() *ArrayInitializerElementContext {
	var p = new(ArrayInitializerElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_arrayInitializerElement
	return p
}

func (*ArrayInitializerElementContext) IsArrayInitializerElementContext() {}

func NewArrayInitializerElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayInitializerElementContext {
	var p = new(ArrayInitializerElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_arrayInitializerElement

	return p
}

func (s *ArrayInitializerElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayInitializerElementContext) MapInitializer() IMapInitializerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapInitializerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapInitializerContext)
}

func (s *ArrayInitializerElementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArrayInitializerElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayInitializerElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayInitializerElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitArrayInitializerElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) ArrayInitializerElement() (localctx IArrayInitializerElementContext) {
	localctx = NewArrayInitializerElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GoScriptParserRULE_arrayInitializerElement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(138)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoScriptParserT__4:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(136)
			p.MapInitializer()
		}

	case GoScriptParserT__2, GoScriptParserT__25, GoScriptParserBooleanLiteral, GoScriptParserIdentifier, GoScriptParserNULL, GoScriptParserHexLiteral, GoScriptParserDecimalLiteral, GoScriptParserOctalLiteral, GoScriptParserFloatingPointLiteral, GoScriptParserPlus, GoScriptParserMinus, GoScriptParserNot, GoScriptParserCharacterLiteral, GoScriptParserStringLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(137)
			p.expression(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMapInitializerContext is an interface to support dynamic dispatch.
type IMapInitializerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapInitializerContext differentiates from other interfaces.
	IsMapInitializerContext()
}

type MapInitializerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapInitializerContext() *MapInitializerContext {
	var p = new(MapInitializerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_mapInitializer
	return p
}

func (*MapInitializerContext) IsMapInitializerContext() {}

func NewMapInitializerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapInitializerContext {
	var p = new(MapInitializerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_mapInitializer

	return p
}

func (s *MapInitializerContext) GetParser() antlr.Parser { return s.parser }

func (s *MapInitializerContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MapInitializerContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MapInitializerContext) AllVariableInitializer() []IVariableInitializerContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableInitializerContext)(nil)).Elem())
	var tst = make([]IVariableInitializerContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableInitializerContext)
		}
	}

	return tst
}

func (s *MapInitializerContext) VariableInitializer(i int) IVariableInitializerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableInitializerContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableInitializerContext)
}

func (s *MapInitializerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapInitializerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapInitializerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitMapInitializer(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) MapInitializer() (localctx IMapInitializerContext) {
	localctx = NewMapInitializerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GoScriptParserRULE_mapInitializer)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		p.Match(GoScriptParserT__4)
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GoScriptParserT__2)|(1<<GoScriptParserT__25)|(1<<GoScriptParserBooleanLiteral))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(GoScriptParserIdentifier-32))|(1<<(GoScriptParserNULL-32))|(1<<(GoScriptParserHexLiteral-32))|(1<<(GoScriptParserDecimalLiteral-32))|(1<<(GoScriptParserOctalLiteral-32))|(1<<(GoScriptParserFloatingPointLiteral-32))|(1<<(GoScriptParserPlus-32))|(1<<(GoScriptParserMinus-32))|(1<<(GoScriptParserNot-32))|(1<<(GoScriptParserCharacterLiteral-32))|(1<<(GoScriptParserStringLiteral-32)))) != 0) {
		{
			p.SetState(141)
			p.expression(0)
		}
		{
			p.SetState(142)
			p.Match(GoScriptParserT__7)
		}
		{
			p.SetState(143)
			p.VariableInitializer()
		}

		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(150)
		p.Match(GoScriptParserT__5)
	}

	return localctx
}

// IType_Context is an interface to support dynamic dispatch.
type IType_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_Context differentiates from other interfaces.
	IsType_Context()
}

type Type_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_Context() *Type_Context {
	var p = new(Type_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_type_
	return p
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) PrimitiveType() IPrimitiveTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitiveTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitiveTypeContext)
}

func (s *Type_Context) Brackets() antlr.TerminalNode {
	return s.GetToken(GoScriptParserBrackets, 0)
}

func (s *Type_Context) MapType() IMapTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapTypeContext)
}

func (s *Type_Context) ConnectorType() IConnectorTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConnectorTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConnectorTypeContext)
}

func (s *Type_Context) DynamicType() IDynamicTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDynamicTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDynamicTypeContext)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitType_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) Type_() (localctx IType_Context) {
	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GoScriptParserRULE_type_)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(168)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoScriptParserT__11, GoScriptParserT__12, GoScriptParserT__13, GoScriptParserT__14, GoScriptParserT__15:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(152)
			p.PrimitiveType()
		}
		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GoScriptParserBrackets {
			{
				p.SetState(153)
				p.Match(GoScriptParserBrackets)
			}

		}

	case GoScriptParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(156)
			p.MapType()
		}
		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GoScriptParserBrackets {
			{
				p.SetState(157)
				p.Match(GoScriptParserBrackets)
			}

		}

	case GoScriptParserT__9:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(160)
			p.ConnectorType()
		}
		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GoScriptParserBrackets {
			{
				p.SetState(161)
				p.Match(GoScriptParserBrackets)
			}

		}

	case GoScriptParserT__10:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(164)
			p.DynamicType()
		}
		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GoScriptParserBrackets {
			{
				p.SetState(165)
				p.Match(GoScriptParserBrackets)
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMapTypeContext is an interface to support dynamic dispatch.
type IMapTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapTypeContext differentiates from other interfaces.
	IsMapTypeContext()
}

type MapTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapTypeContext() *MapTypeContext {
	var p = new(MapTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_mapType
	return p
}

func (*MapTypeContext) IsMapTypeContext() {}

func NewMapTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapTypeContext {
	var p = new(MapTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_mapType

	return p
}

func (s *MapTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MapTypeContext) Lt() antlr.TerminalNode {
	return s.GetToken(GoScriptParserLt, 0)
}

func (s *MapTypeContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *MapTypeContext) Gt() antlr.TerminalNode {
	return s.GetToken(GoScriptParserGt, 0)
}

func (s *MapTypeContext) PrimitiveType() IPrimitiveTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitiveTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitiveTypeContext)
}

func (s *MapTypeContext) ConnectorType() IConnectorTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConnectorTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConnectorTypeContext)
}

func (s *MapTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitMapType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) MapType() (localctx IMapTypeContext) {
	localctx = NewMapTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GoScriptParserRULE_mapType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Match(GoScriptParserT__8)
	}
	{
		p.SetState(171)
		p.Match(GoScriptParserLt)
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoScriptParserT__11, GoScriptParserT__12, GoScriptParserT__13, GoScriptParserT__14, GoScriptParserT__15:
		{
			p.SetState(172)
			p.PrimitiveType()
		}

	case GoScriptParserT__9:
		{
			p.SetState(173)
			p.ConnectorType()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(176)
		p.Match(GoScriptParserT__6)
	}
	{
		p.SetState(177)
		p.Type_()
	}
	{
		p.SetState(178)
		p.Match(GoScriptParserGt)
	}

	return localctx
}

// IConnectorTypeContext is an interface to support dynamic dispatch.
type IConnectorTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConnectorTypeContext differentiates from other interfaces.
	IsConnectorTypeContext()
}

type ConnectorTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConnectorTypeContext() *ConnectorTypeContext {
	var p = new(ConnectorTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_connectorType
	return p
}

func (*ConnectorTypeContext) IsConnectorTypeContext() {}

func NewConnectorTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConnectorTypeContext {
	var p = new(ConnectorTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_connectorType

	return p
}

func (s *ConnectorTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ConnectorTypeContext) Lt() antlr.TerminalNode {
	return s.GetToken(GoScriptParserLt, 0)
}

func (s *ConnectorTypeContext) Identifier() antlr.TerminalNode {
	return s.GetToken(GoScriptParserIdentifier, 0)
}

func (s *ConnectorTypeContext) Gt() antlr.TerminalNode {
	return s.GetToken(GoScriptParserGt, 0)
}

func (s *ConnectorTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConnectorTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConnectorTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitConnectorType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) ConnectorType() (localctx IConnectorTypeContext) {
	localctx = NewConnectorTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GoScriptParserRULE_connectorType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.Match(GoScriptParserT__9)
	}
	{
		p.SetState(181)
		p.Match(GoScriptParserLt)
	}
	{
		p.SetState(182)
		p.Match(GoScriptParserIdentifier)
	}
	{
		p.SetState(183)
		p.Match(GoScriptParserGt)
	}

	return localctx
}

// IDynamicTypeContext is an interface to support dynamic dispatch.
type IDynamicTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDynamicTypeContext differentiates from other interfaces.
	IsDynamicTypeContext()
}

type DynamicTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDynamicTypeContext() *DynamicTypeContext {
	var p = new(DynamicTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_dynamicType
	return p
}

func (*DynamicTypeContext) IsDynamicTypeContext() {}

func NewDynamicTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DynamicTypeContext {
	var p = new(DynamicTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_dynamicType

	return p
}

func (s *DynamicTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *DynamicTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DynamicTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DynamicTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitDynamicType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) DynamicType() (localctx IDynamicTypeContext) {
	localctx = NewDynamicTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GoScriptParserRULE_dynamicType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(185)
		p.Match(GoScriptParserT__10)
	}

	return localctx
}

// IPrimitiveTypeContext is an interface to support dynamic dispatch.
type IPrimitiveTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimitiveTypeContext differentiates from other interfaces.
	IsPrimitiveTypeContext()
}

type PrimitiveTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitiveTypeContext() *PrimitiveTypeContext {
	var p = new(PrimitiveTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_primitiveType
	return p
}

func (*PrimitiveTypeContext) IsPrimitiveTypeContext() {}

func NewPrimitiveTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveTypeContext {
	var p = new(PrimitiveTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_primitiveType

	return p
}

func (s *PrimitiveTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *PrimitiveTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitiveTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitPrimitiveType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) PrimitiveType() (localctx IPrimitiveTypeContext) {
	localctx = NewPrimitiveTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GoScriptParserRULE_primitiveType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GoScriptParserT__11)|(1<<GoScriptParserT__12)|(1<<GoScriptParserT__13)|(1<<GoScriptParserT__14)|(1<<GoScriptParserT__15))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyFrom(ctx *StatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BreakStatementContext struct {
	*StatementContext
}

func NewBreakStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakStatementContext {
	var p = new(BreakStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *BreakStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitBreakStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfStatementContext struct {
	*StatementContext
}

func NewIfStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStatementContext {
	var p = new(IfStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *IfStatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type CommonBlockStatementContext struct {
	*StatementContext
}

func NewCommonBlockStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CommonBlockStatementContext {
	var p = new(CommonBlockStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *CommonBlockStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommonBlockStatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *CommonBlockStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitCommonBlockStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReturnStatementContext struct {
	*StatementContext
}

func NewReturnStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprStatementContext struct {
	*StatementContext
}

func NewExprStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprStatementContext {
	var p = new(ExprStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *ExprStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExprStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitExprStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForStatementContext struct {
	*StatementContext
}

func NewForStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForStatementContext {
	var p = new(ForStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *ForStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatementContext) ForControl() IForControlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForControlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForControlContext)
}

func (s *ForStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ForStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitForStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ContinueStatementContext struct {
	*StatementContext
}

func NewContinueStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinueStatementContext {
	var p = new(ContinueStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *ContinueStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitContinueStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GoScriptParserRULE_statement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(215)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoScriptParserT__4:
		localctx = NewCommonBlockStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(189)
			p.Block()
		}

	case GoScriptParserT__16:
		localctx = NewIfStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(190)
			p.Match(GoScriptParserT__16)
		}
		{
			p.SetState(191)
			p.expression(0)
		}
		{
			p.SetState(192)
			p.Statement()
		}
		p.SetState(195)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(193)
				p.Match(GoScriptParserT__17)
			}
			{
				p.SetState(194)
				p.Statement()
			}

		}

	case GoScriptParserT__18:
		localctx = NewForStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(197)
			p.Match(GoScriptParserT__18)
		}
		{
			p.SetState(198)
			p.Match(GoScriptParserT__2)
		}
		{
			p.SetState(199)
			p.ForControl()
		}
		{
			p.SetState(200)
			p.Match(GoScriptParserT__3)
		}
		{
			p.SetState(201)
			p.Statement()
		}

	case GoScriptParserT__19:
		localctx = NewReturnStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(203)
			p.Match(GoScriptParserT__19)
		}
		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GoScriptParserT__2)|(1<<GoScriptParserT__25)|(1<<GoScriptParserBooleanLiteral))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(GoScriptParserIdentifier-32))|(1<<(GoScriptParserNULL-32))|(1<<(GoScriptParserHexLiteral-32))|(1<<(GoScriptParserDecimalLiteral-32))|(1<<(GoScriptParserOctalLiteral-32))|(1<<(GoScriptParserFloatingPointLiteral-32))|(1<<(GoScriptParserPlus-32))|(1<<(GoScriptParserMinus-32))|(1<<(GoScriptParserNot-32))|(1<<(GoScriptParserCharacterLiteral-32))|(1<<(GoScriptParserStringLiteral-32)))) != 0) {
			{
				p.SetState(204)
				p.expression(0)
			}

		}
		{
			p.SetState(207)
			p.Match(GoScriptParserT__0)
		}

	case GoScriptParserT__20:
		localctx = NewBreakStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(208)
			p.Match(GoScriptParserT__20)
		}
		{
			p.SetState(209)
			p.Match(GoScriptParserT__0)
		}

	case GoScriptParserT__21:
		localctx = NewContinueStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(210)
			p.Match(GoScriptParserT__21)
		}
		{
			p.SetState(211)
			p.Match(GoScriptParserT__0)
		}

	case GoScriptParserT__2, GoScriptParserT__25, GoScriptParserBooleanLiteral, GoScriptParserIdentifier, GoScriptParserNULL, GoScriptParserHexLiteral, GoScriptParserDecimalLiteral, GoScriptParserOctalLiteral, GoScriptParserFloatingPointLiteral, GoScriptParserPlus, GoScriptParserMinus, GoScriptParserNot, GoScriptParserCharacterLiteral, GoScriptParserStringLiteral:
		localctx = NewExprStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(212)
			p.expression(0)
		}
		{
			p.SetState(213)
			p.Match(GoScriptParserT__0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IForControlContext is an interface to support dynamic dispatch.
type IForControlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForControlContext differentiates from other interfaces.
	IsForControlContext()
}

type ForControlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForControlContext() *ForControlContext {
	var p = new(ForControlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_forControl
	return p
}

func (*ForControlContext) IsForControlContext() {}

func NewForControlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForControlContext {
	var p = new(ForControlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_forControl

	return p
}

func (s *ForControlContext) GetParser() antlr.Parser { return s.parser }

func (s *ForControlContext) ForInit() IForInitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForInitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForInitContext)
}

func (s *ForControlContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ForControlContext) ForUpdate() IForUpdateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForUpdateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForUpdateContext)
}

func (s *ForControlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForControlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForControlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitForControl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) ForControl() (localctx IForControlContext) {
	localctx = NewForControlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GoScriptParserRULE_forControl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GoScriptParserT__2)|(1<<GoScriptParserT__8)|(1<<GoScriptParserT__9)|(1<<GoScriptParserT__10)|(1<<GoScriptParserT__11)|(1<<GoScriptParserT__12)|(1<<GoScriptParserT__13)|(1<<GoScriptParserT__14)|(1<<GoScriptParserT__15)|(1<<GoScriptParserT__25)|(1<<GoScriptParserBooleanLiteral))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(GoScriptParserIdentifier-32))|(1<<(GoScriptParserNULL-32))|(1<<(GoScriptParserHexLiteral-32))|(1<<(GoScriptParserDecimalLiteral-32))|(1<<(GoScriptParserOctalLiteral-32))|(1<<(GoScriptParserFloatingPointLiteral-32))|(1<<(GoScriptParserPlus-32))|(1<<(GoScriptParserMinus-32))|(1<<(GoScriptParserNot-32))|(1<<(GoScriptParserCharacterLiteral-32))|(1<<(GoScriptParserStringLiteral-32)))) != 0) {
		{
			p.SetState(217)
			p.ForInit()
		}

	}
	{
		p.SetState(220)
		p.Match(GoScriptParserT__0)
	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GoScriptParserT__2)|(1<<GoScriptParserT__25)|(1<<GoScriptParserBooleanLiteral))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(GoScriptParserIdentifier-32))|(1<<(GoScriptParserNULL-32))|(1<<(GoScriptParserHexLiteral-32))|(1<<(GoScriptParserDecimalLiteral-32))|(1<<(GoScriptParserOctalLiteral-32))|(1<<(GoScriptParserFloatingPointLiteral-32))|(1<<(GoScriptParserPlus-32))|(1<<(GoScriptParserMinus-32))|(1<<(GoScriptParserNot-32))|(1<<(GoScriptParserCharacterLiteral-32))|(1<<(GoScriptParserStringLiteral-32)))) != 0) {
		{
			p.SetState(221)
			p.expression(0)
		}

	}
	{
		p.SetState(224)
		p.Match(GoScriptParserT__0)
	}
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GoScriptParserT__2)|(1<<GoScriptParserT__25)|(1<<GoScriptParserBooleanLiteral))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(GoScriptParserIdentifier-32))|(1<<(GoScriptParserNULL-32))|(1<<(GoScriptParserHexLiteral-32))|(1<<(GoScriptParserDecimalLiteral-32))|(1<<(GoScriptParserOctalLiteral-32))|(1<<(GoScriptParserFloatingPointLiteral-32))|(1<<(GoScriptParserPlus-32))|(1<<(GoScriptParserMinus-32))|(1<<(GoScriptParserNot-32))|(1<<(GoScriptParserCharacterLiteral-32))|(1<<(GoScriptParserStringLiteral-32)))) != 0) {
		{
			p.SetState(225)
			p.ForUpdate()
		}

	}

	return localctx
}

// IForInitContext is an interface to support dynamic dispatch.
type IForInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForInitContext differentiates from other interfaces.
	IsForInitContext()
}

type ForInitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForInitContext() *ForInitContext {
	var p = new(ForInitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_forInit
	return p
}

func (*ForInitContext) IsForInitContext() {}

func NewForInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForInitContext {
	var p = new(ForInitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_forInit

	return p
}

func (s *ForInitContext) GetParser() antlr.Parser { return s.parser }

func (s *ForInitContext) VariableDeclaration() IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *ForInitContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ForInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForInitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitForInit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) ForInit() (localctx IForInitContext) {
	localctx = NewForInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GoScriptParserRULE_forInit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(230)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoScriptParserT__8, GoScriptParserT__9, GoScriptParserT__10, GoScriptParserT__11, GoScriptParserT__12, GoScriptParserT__13, GoScriptParserT__14, GoScriptParserT__15:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(228)
			p.VariableDeclaration()
		}

	case GoScriptParserT__2, GoScriptParserT__25, GoScriptParserBooleanLiteral, GoScriptParserIdentifier, GoScriptParserNULL, GoScriptParserHexLiteral, GoScriptParserDecimalLiteral, GoScriptParserOctalLiteral, GoScriptParserFloatingPointLiteral, GoScriptParserPlus, GoScriptParserMinus, GoScriptParserNot, GoScriptParserCharacterLiteral, GoScriptParserStringLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(229)
			p.ExpressionList()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IForUpdateContext is an interface to support dynamic dispatch.
type IForUpdateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForUpdateContext differentiates from other interfaces.
	IsForUpdateContext()
}

type ForUpdateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForUpdateContext() *ForUpdateContext {
	var p = new(ForUpdateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_forUpdate
	return p
}

func (*ForUpdateContext) IsForUpdateContext() {}

func NewForUpdateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForUpdateContext {
	var p = new(ForUpdateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_forUpdate

	return p
}

func (s *ForUpdateContext) GetParser() antlr.Parser { return s.parser }

func (s *ForUpdateContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ForUpdateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForUpdateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForUpdateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitForUpdate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) ForUpdate() (localctx IForUpdateContext) {
	localctx = NewForUpdateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GoScriptParserRULE_forUpdate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(232)
		p.ExpressionList()
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AndExprContext struct {
	*ExpressionContext
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AndExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CreatorExprContext struct {
	*ExpressionContext
}

func NewCreatorExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CreatorExprContext {
	var p = new(CreatorExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *CreatorExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreatorExprContext) Creator() ICreatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICreatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICreatorContext)
}

func (s *CreatorExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitCreatorExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type FetchExprContext struct {
	*ExpressionContext
}

func NewFetchExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FetchExprContext {
	var p = new(FetchExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *FetchExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FetchExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *FetchExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FetchExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitFetchExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrExprContext struct {
	*ExpressionContext
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *OrExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IndexExprContext struct {
	*ExpressionContext
}

func NewIndexExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexExprContext {
	var p = new(IndexExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IndexExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *IndexExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IndexExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitIndexExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CreatorAndAssignExprContext struct {
	*ExpressionContext
}

func NewCreatorAndAssignExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CreatorAndAssignExprContext {
	var p = new(CreatorAndAssignExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *CreatorAndAssignExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreatorAndAssignExprContext) Identifier() antlr.TerminalNode {
	return s.GetToken(GoScriptParserIdentifier, 0)
}

func (s *CreatorAndAssignExprContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CreatorAndAssignExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitCreatorAndAssignExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignExprContext struct {
	*ExpressionContext
}

func NewAssignExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignExprContext {
	var p = new(AssignExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AssignExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AssignExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignExprContext) Assign() antlr.TerminalNode {
	return s.GetToken(GoScriptParserAssign, 0)
}

func (s *AssignExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitAssignExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NegativeExprContext struct {
	*ExpressionContext
}

func NewNegativeExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegativeExprContext {
	var p = new(NegativeExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NegativeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegativeExprContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NegativeExprContext) Plus() antlr.TerminalNode {
	return s.GetToken(GoScriptParserPlus, 0)
}

func (s *NegativeExprContext) Minus() antlr.TerminalNode {
	return s.GetToken(GoScriptParserMinus, 0)
}

func (s *NegativeExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitNegativeExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulDivExprContext struct {
	*ExpressionContext
}

func NewMulDivExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivExprContext {
	var p = new(MulDivExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MulDivExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MulDivExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MulDivExprContext) Multiplication() antlr.TerminalNode {
	return s.GetToken(GoScriptParserMultiplication, 0)
}

func (s *MulDivExprContext) Division() antlr.TerminalNode {
	return s.GetToken(GoScriptParserDivision, 0)
}

func (s *MulDivExprContext) Percent() antlr.TerminalNode {
	return s.GetToken(GoScriptParserPercent, 0)
}

func (s *MulDivExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitMulDivExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CompareExprContext struct {
	*ExpressionContext
}

func NewCompareExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareExprContext {
	var p = new(CompareExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *CompareExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *CompareExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CompareExprContext) Lt() antlr.TerminalNode {
	return s.GetToken(GoScriptParserLt, 0)
}

func (s *CompareExprContext) Assign() antlr.TerminalNode {
	return s.GetToken(GoScriptParserAssign, 0)
}

func (s *CompareExprContext) Gt() antlr.TerminalNode {
	return s.GetToken(GoScriptParserGt, 0)
}

func (s *CompareExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitCompareExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type EqNEqExprContext struct {
	*ExpressionContext
}

func NewEqNEqExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqNEqExprContext {
	var p = new(EqNEqExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *EqNEqExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqNEqExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *EqNEqExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EqNEqExprContext) Eq() antlr.TerminalNode {
	return s.GetToken(GoScriptParserEq, 0)
}

func (s *EqNEqExprContext) NEq() antlr.TerminalNode {
	return s.GetToken(GoScriptParserNEq, 0)
}

func (s *EqNEqExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitEqNEqExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrimaryExprContext struct {
	*ExpressionContext
}

func NewPrimaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimaryExprContext {
	var p = new(PrimaryExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExprContext) Primary() IPrimaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *PrimaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitPrimaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CallExprContext struct {
	*ExpressionContext
}

func NewCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallExprContext {
	var p = new(CallExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *CallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallExprContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CallExprContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *CallExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitCallExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotExprContext struct {
	*ExpressionContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) Not() antlr.TerminalNode {
	return s.GetToken(GoScriptParserNot, 0)
}

func (s *NotExprContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitNotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddSubExprContext struct {
	*ExpressionContext
}

func NewAddSubExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubExprContext {
	var p = new(AddSubExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AddSubExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AddSubExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddSubExprContext) Plus() antlr.TerminalNode {
	return s.GetToken(GoScriptParserPlus, 0)
}

func (s *AddSubExprContext) Minus() antlr.TerminalNode {
	return s.GetToken(GoScriptParserMinus, 0)
}

func (s *AddSubExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitAddSubExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type TernaryExprContext struct {
	*ExpressionContext
}

func NewTernaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TernaryExprContext {
	var p = new(TernaryExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *TernaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TernaryExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *TernaryExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TernaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitTernaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *GoScriptParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 44
	p.EnterRecursionRule(localctx, 44, GoScriptParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPrimaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(235)
			p.Primary()
		}

	case 2:
		localctx = NewNegativeExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(236)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GoScriptParserPlus || _la == GoScriptParserMinus) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(237)
			p.expression(12)
		}

	case 3:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(238)
			p.Match(GoScriptParserNot)
		}
		{
			p.SetState(239)
			p.expression(11)
		}

	case 4:
		localctx = NewCreatorExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(240)
			p.Match(GoScriptParserT__25)
		}
		{
			p.SetState(241)
			p.Creator()
		}

	case 5:
		localctx = NewCreatorAndAssignExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(242)
			p.Match(GoScriptParserIdentifier)
		}
		{
			p.SetState(243)
			p.Match(GoScriptParserT__29)
		}
		{
			p.SetState(244)
			p.expression(1)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(295)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
			case 1:
				localctx = NewFetchExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(247)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(248)
					p.Match(GoScriptParserT__22)
				}
				{
					p.SetState(249)
					p.expression(16)
				}

			case 2:
				localctx = NewMulDivExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(250)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(251)
					_la = p.GetTokenStream().LA(1)

					if !(((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(GoScriptParserMultiplication-41))|(1<<(GoScriptParserDivision-41))|(1<<(GoScriptParserPercent-41)))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(252)
					p.expression(10)
				}

			case 3:
				localctx = NewAddSubExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(253)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(254)
					_la = p.GetTokenStream().LA(1)

					if !(_la == GoScriptParserPlus || _la == GoScriptParserMinus) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(255)
					p.expression(9)
				}

			case 4:
				localctx = NewCompareExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(256)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				p.SetState(263)
				p.GetErrorHandler().Sync(p)
				switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
				case 1:
					{
						p.SetState(257)
						p.Match(GoScriptParserLt)
					}
					{
						p.SetState(258)
						p.Match(GoScriptParserAssign)
					}

				case 2:
					{
						p.SetState(259)
						p.Match(GoScriptParserGt)
					}
					{
						p.SetState(260)
						p.Match(GoScriptParserAssign)
					}

				case 3:
					{
						p.SetState(261)
						p.Match(GoScriptParserGt)
					}

				case 4:
					{
						p.SetState(262)
						p.Match(GoScriptParserLt)
					}

				}
				{
					p.SetState(265)
					p.expression(8)
				}

			case 5:
				localctx = NewEqNEqExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(266)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(267)
					_la = p.GetTokenStream().LA(1)

					if !(_la == GoScriptParserEq || _la == GoScriptParserNEq) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(268)
					p.expression(7)
				}

			case 6:
				localctx = NewAndExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(269)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(270)
					p.Match(GoScriptParserT__26)
				}
				{
					p.SetState(271)
					p.expression(6)
				}

			case 7:
				localctx = NewOrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(272)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(273)
					p.Match(GoScriptParserT__27)
				}
				{
					p.SetState(274)
					p.expression(5)
				}

			case 8:
				localctx = NewTernaryExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(275)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(276)
					p.Match(GoScriptParserT__28)
				}
				{
					p.SetState(277)
					p.expression(0)
				}
				{
					p.SetState(278)
					p.Match(GoScriptParserT__7)
				}
				{
					p.SetState(279)
					p.expression(4)
				}

			case 9:
				localctx = NewAssignExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(281)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(282)
					p.Match(GoScriptParserAssign)
				}
				{
					p.SetState(283)
					p.expression(2)
				}

			case 10:
				localctx = NewIndexExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(284)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(285)
					p.Match(GoScriptParserT__23)
				}
				{
					p.SetState(286)
					p.expression(0)
				}
				{
					p.SetState(287)
					p.Match(GoScriptParserT__24)
				}

			case 11:
				localctx = NewCallExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(289)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(290)
					p.Match(GoScriptParserT__2)
				}
				p.SetState(292)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GoScriptParserT__2)|(1<<GoScriptParserT__25)|(1<<GoScriptParserBooleanLiteral))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(GoScriptParserIdentifier-32))|(1<<(GoScriptParserNULL-32))|(1<<(GoScriptParserHexLiteral-32))|(1<<(GoScriptParserDecimalLiteral-32))|(1<<(GoScriptParserOctalLiteral-32))|(1<<(GoScriptParserFloatingPointLiteral-32))|(1<<(GoScriptParserPlus-32))|(1<<(GoScriptParserMinus-32))|(1<<(GoScriptParserNot-32))|(1<<(GoScriptParserCharacterLiteral-32))|(1<<(GoScriptParserStringLiteral-32)))) != 0) {
					{
						p.SetState(291)
						p.ExpressionList()
					}

				}
				{
					p.SetState(294)
					p.Match(GoScriptParserT__3)
				}

			}

		}
		p.SetState(299)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_primary
	return p
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrimaryContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *PrimaryContext) Identifier() antlr.TerminalNode {
	return s.GetToken(GoScriptParserIdentifier, 0)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, GoScriptParserRULE_primary)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(306)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoScriptParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(300)
			p.Match(GoScriptParserT__2)
		}
		{
			p.SetState(301)
			p.expression(0)
		}
		{
			p.SetState(302)
			p.Match(GoScriptParserT__3)
		}

	case GoScriptParserBooleanLiteral, GoScriptParserNULL, GoScriptParserHexLiteral, GoScriptParserDecimalLiteral, GoScriptParserOctalLiteral, GoScriptParserFloatingPointLiteral, GoScriptParserCharacterLiteral, GoScriptParserStringLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(304)
			p.Literal()
		}

	case GoScriptParserIdentifier:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(305)
			p.Match(GoScriptParserIdentifier)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) IntegerLiteral() IIntegerLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerLiteralContext)
}

func (s *LiteralContext) FloatingPointLiteral() antlr.TerminalNode {
	return s.GetToken(GoScriptParserFloatingPointLiteral, 0)
}

func (s *LiteralContext) CharacterLiteral() antlr.TerminalNode {
	return s.GetToken(GoScriptParserCharacterLiteral, 0)
}

func (s *LiteralContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(GoScriptParserStringLiteral, 0)
}

func (s *LiteralContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(GoScriptParserBooleanLiteral, 0)
}

func (s *LiteralContext) NULL() antlr.TerminalNode {
	return s.GetToken(GoScriptParserNULL, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, GoScriptParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(314)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoScriptParserHexLiteral, GoScriptParserDecimalLiteral, GoScriptParserOctalLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(308)
			p.IntegerLiteral()
		}

	case GoScriptParserFloatingPointLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(309)
			p.Match(GoScriptParserFloatingPointLiteral)
		}

	case GoScriptParserCharacterLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(310)
			p.Match(GoScriptParserCharacterLiteral)
		}

	case GoScriptParserStringLiteral:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(311)
			p.Match(GoScriptParserStringLiteral)
		}

	case GoScriptParserBooleanLiteral:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(312)
			p.Match(GoScriptParserBooleanLiteral)
		}

	case GoScriptParserNULL:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(313)
			p.Match(GoScriptParserNULL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIntegerLiteralContext is an interface to support dynamic dispatch.
type IIntegerLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerLiteralContext differentiates from other interfaces.
	IsIntegerLiteralContext()
}

type IntegerLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerLiteralContext() *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_integerLiteral
	return p
}

func (*IntegerLiteralContext) IsIntegerLiteralContext() {}

func NewIntegerLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_integerLiteral

	return p
}

func (s *IntegerLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerLiteralContext) HexLiteral() antlr.TerminalNode {
	return s.GetToken(GoScriptParserHexLiteral, 0)
}

func (s *IntegerLiteralContext) OctalLiteral() antlr.TerminalNode {
	return s.GetToken(GoScriptParserOctalLiteral, 0)
}

func (s *IntegerLiteralContext) DecimalLiteral() antlr.TerminalNode {
	return s.GetToken(GoScriptParserDecimalLiteral, 0)
}

func (s *IntegerLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitIntegerLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) IntegerLiteral() (localctx IIntegerLiteralContext) {
	localctx = NewIntegerLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, GoScriptParserRULE_integerLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(316)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(GoScriptParserHexLiteral-35))|(1<<(GoScriptParserDecimalLiteral-35))|(1<<(GoScriptParserOctalLiteral-35)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_expressionList
	return p
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, GoScriptParserRULE_expressionList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(318)
		p.expression(0)
	}
	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoScriptParserT__6 {
		{
			p.SetState(319)
			p.Match(GoScriptParserT__6)
		}
		{
			p.SetState(320)
			p.expression(0)
		}

		p.SetState(325)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICreatorContext is an interface to support dynamic dispatch.
type ICreatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreatorContext differentiates from other interfaces.
	IsCreatorContext()
}

type CreatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreatorContext() *CreatorContext {
	var p = new(CreatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_creator
	return p
}

func (*CreatorContext) IsCreatorContext() {}

func NewCreatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreatorContext {
	var p = new(CreatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_creator

	return p
}

func (s *CreatorContext) GetParser() antlr.Parser { return s.parser }

func (s *CreatorContext) CopyFrom(ctx *CreatorContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *CreatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MeegoCreatorContext struct {
	*CreatorContext
}

func NewMeegoCreatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MeegoCreatorContext {
	var p = new(MeegoCreatorContext)

	p.CreatorContext = NewEmptyCreatorContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CreatorContext))

	return p
}

func (s *MeegoCreatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MeegoCreatorContext) ConnectorType() IConnectorTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConnectorTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConnectorTypeContext)
}

func (s *MeegoCreatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitMeegoCreator(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrimitiveCreatorContext struct {
	*CreatorContext
}

func NewPrimitiveCreatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimitiveCreatorContext {
	var p = new(PrimitiveCreatorContext)

	p.CreatorContext = NewEmptyCreatorContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CreatorContext))

	return p
}

func (s *PrimitiveCreatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveCreatorContext) PrimitiveType() IPrimitiveTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitiveTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitiveTypeContext)
}

func (s *PrimitiveCreatorContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrimitiveCreatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitPrimitiveCreator(s)

	default:
		return t.VisitChildren(s)
	}
}

type ArrayCreatorContext struct {
	*CreatorContext
}

func NewArrayCreatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayCreatorContext {
	var p = new(ArrayCreatorContext)

	p.CreatorContext = NewEmptyCreatorContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CreatorContext))

	return p
}

func (s *ArrayCreatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayCreatorContext) CreatorName() ICreatorNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICreatorNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICreatorNameContext)
}

func (s *ArrayCreatorContext) ArrayInitializer() IArrayInitializerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayInitializerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayInitializerContext)
}

func (s *ArrayCreatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitArrayCreator(s)

	default:
		return t.VisitChildren(s)
	}
}

type DynamicCreatorContext struct {
	*CreatorContext
}

func NewDynamicCreatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DynamicCreatorContext {
	var p = new(DynamicCreatorContext)

	p.CreatorContext = NewEmptyCreatorContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CreatorContext))

	return p
}

func (s *DynamicCreatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DynamicCreatorContext) DynamicType() IDynamicTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDynamicTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDynamicTypeContext)
}

func (s *DynamicCreatorContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DynamicCreatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitDynamicCreator(s)

	default:
		return t.VisitChildren(s)
	}
}

type MapCreatorContext struct {
	*CreatorContext
}

func NewMapCreatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MapCreatorContext {
	var p = new(MapCreatorContext)

	p.CreatorContext = NewEmptyCreatorContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CreatorContext))

	return p
}

func (s *MapCreatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapCreatorContext) MapType() IMapTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapTypeContext)
}

func (s *MapCreatorContext) MapInitializer() IMapInitializerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapInitializerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapInitializerContext)
}

func (s *MapCreatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitMapCreator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) Creator() (localctx ICreatorContext) {
	localctx = NewCreatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, GoScriptParserRULE_creator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		localctx = NewMapCreatorContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(326)
			p.MapType()
		}
		{
			p.SetState(327)
			p.MapInitializer()
		}

	case 2:
		localctx = NewArrayCreatorContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(329)
			p.CreatorName()
		}
		{
			p.SetState(330)
			p.Match(GoScriptParserT__23)
		}
		{
			p.SetState(331)
			p.Match(GoScriptParserT__24)
		}
		{
			p.SetState(332)
			p.ArrayInitializer()
		}

	case 3:
		localctx = NewMeegoCreatorContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(334)
			p.ConnectorType()
		}
		{
			p.SetState(335)
			p.Match(GoScriptParserT__2)
		}
		{
			p.SetState(336)
			p.Match(GoScriptParserT__3)
		}

	case 4:
		localctx = NewPrimitiveCreatorContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(338)
			p.PrimitiveType()
		}
		{
			p.SetState(339)
			p.Match(GoScriptParserT__2)
		}
		p.SetState(341)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GoScriptParserT__2)|(1<<GoScriptParserT__25)|(1<<GoScriptParserBooleanLiteral))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(GoScriptParserIdentifier-32))|(1<<(GoScriptParserNULL-32))|(1<<(GoScriptParserHexLiteral-32))|(1<<(GoScriptParserDecimalLiteral-32))|(1<<(GoScriptParserOctalLiteral-32))|(1<<(GoScriptParserFloatingPointLiteral-32))|(1<<(GoScriptParserPlus-32))|(1<<(GoScriptParserMinus-32))|(1<<(GoScriptParserNot-32))|(1<<(GoScriptParserCharacterLiteral-32))|(1<<(GoScriptParserStringLiteral-32)))) != 0) {
			{
				p.SetState(340)
				p.expression(0)
			}

		}
		{
			p.SetState(343)
			p.Match(GoScriptParserT__3)
		}

	case 5:
		localctx = NewDynamicCreatorContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(345)
			p.DynamicType()
		}
		{
			p.SetState(346)
			p.Match(GoScriptParserT__2)
		}
		p.SetState(348)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GoScriptParserT__2)|(1<<GoScriptParserT__25)|(1<<GoScriptParserBooleanLiteral))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(GoScriptParserIdentifier-32))|(1<<(GoScriptParserNULL-32))|(1<<(GoScriptParserHexLiteral-32))|(1<<(GoScriptParserDecimalLiteral-32))|(1<<(GoScriptParserOctalLiteral-32))|(1<<(GoScriptParserFloatingPointLiteral-32))|(1<<(GoScriptParserPlus-32))|(1<<(GoScriptParserMinus-32))|(1<<(GoScriptParserNot-32))|(1<<(GoScriptParserCharacterLiteral-32))|(1<<(GoScriptParserStringLiteral-32)))) != 0) {
			{
				p.SetState(347)
				p.expression(0)
			}

		}
		{
			p.SetState(350)
			p.Match(GoScriptParserT__3)
		}

	}

	return localctx
}

// ICreatorNameContext is an interface to support dynamic dispatch.
type ICreatorNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreatorNameContext differentiates from other interfaces.
	IsCreatorNameContext()
}

type CreatorNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreatorNameContext() *CreatorNameContext {
	var p = new(CreatorNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_creatorName
	return p
}

func (*CreatorNameContext) IsCreatorNameContext() {}

func NewCreatorNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreatorNameContext {
	var p = new(CreatorNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_creatorName

	return p
}

func (s *CreatorNameContext) GetParser() antlr.Parser { return s.parser }

func (s *CreatorNameContext) PrimitiveType() IPrimitiveTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitiveTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitiveTypeContext)
}

func (s *CreatorNameContext) MapType() IMapTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapTypeContext)
}

func (s *CreatorNameContext) ConnectorType() IConnectorTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConnectorTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConnectorTypeContext)
}

func (s *CreatorNameContext) DynamicType() IDynamicTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDynamicTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDynamicTypeContext)
}

func (s *CreatorNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreatorNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreatorNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitCreatorName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) CreatorName() (localctx ICreatorNameContext) {
	localctx = NewCreatorNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, GoScriptParserRULE_creatorName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(358)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoScriptParserT__11, GoScriptParserT__12, GoScriptParserT__13, GoScriptParserT__14, GoScriptParserT__15:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(354)
			p.PrimitiveType()
		}

	case GoScriptParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(355)
			p.MapType()
		}

	case GoScriptParserT__9:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(356)
			p.ConnectorType()
		}

	case GoScriptParserT__10:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(357)
			p.DynamicType()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *GoScriptParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 22:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GoScriptParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 13)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
