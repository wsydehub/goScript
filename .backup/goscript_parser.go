// Code generated from GoScript.g4 by ANTLR 4.9.2. DO NOT EDIT.

package goScript // GoScript

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 62, 427,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 85, 10, 2, 12, 2, 14, 2,
	88, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 96, 10, 3, 12, 3,
	14, 3, 99, 11, 3, 3, 3, 3, 3, 3, 4, 3, 4, 7, 4, 105, 10, 4, 12, 4, 14,
	4, 108, 11, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 7, 6, 117, 10,
	6, 12, 6, 14, 6, 120, 11, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7,
	128, 10, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 7, 9, 136, 10, 9, 12, 9,
	14, 9, 139, 11, 9, 3, 10, 3, 10, 3, 10, 5, 10, 144, 10, 10, 3, 11, 3, 11,
	3, 11, 5, 11, 149, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 155, 10,
	12, 12, 12, 14, 12, 158, 11, 12, 5, 12, 160, 10, 12, 3, 12, 3, 12, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 169, 10, 13, 12, 13, 14, 13, 172, 11,
	13, 3, 13, 3, 13, 3, 14, 3, 14, 7, 14, 178, 10, 14, 12, 14, 14, 14, 181,
	11, 14, 3, 14, 3, 14, 7, 14, 185, 10, 14, 12, 14, 14, 14, 188, 11, 14,
	3, 14, 3, 14, 7, 14, 192, 10, 14, 12, 14, 14, 14, 195, 11, 14, 3, 14, 3,
	14, 7, 14, 199, 10, 14, 12, 14, 14, 14, 202, 11, 14, 5, 14, 204, 10, 14,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 5, 19, 229, 10, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5,
	20, 236, 10, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 5, 22,
	245, 10, 22, 3, 22, 3, 22, 5, 22, 249, 10, 22, 3, 22, 3, 22, 5, 22, 253,
	10, 22, 3, 23, 3, 23, 5, 23, 257, 10, 23, 3, 24, 3, 24, 3, 25, 3, 25, 5,
	25, 263, 10, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27,
	3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 5, 29, 282,
	10, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29,
	3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3,
	29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29,
	3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 7, 29, 323,
	10, 29, 12, 29, 14, 29, 326, 11, 29, 3, 29, 3, 29, 3, 29, 3, 29, 7, 29,
	332, 10, 29, 12, 29, 14, 29, 335, 11, 29, 3, 29, 3, 29, 3, 29, 7, 29, 340,
	10, 29, 12, 29, 14, 29, 343, 11, 29, 3, 29, 3, 29, 3, 29, 3, 29, 7, 29,
	349, 10, 29, 12, 29, 14, 29, 352, 11, 29, 7, 29, 354, 10, 29, 12, 29, 14,
	29, 357, 11, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30, 365,
	10, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 5, 31, 373, 10, 31, 3,
	32, 3, 32, 3, 33, 3, 33, 3, 33, 7, 33, 380, 10, 33, 12, 33, 14, 33, 383,
	11, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 5, 34, 390, 10, 34, 3, 35, 3,
	35, 3, 35, 3, 36, 3, 36, 6, 36, 397, 10, 36, 13, 36, 14, 36, 398, 3, 36,
	3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 5, 37, 407, 10, 37, 3, 38, 3, 38, 3,
	38, 3, 38, 3, 39, 3, 39, 3, 39, 5, 39, 416, 10, 39, 3, 39, 3, 39, 3, 40,
	3, 40, 3, 40, 5, 40, 423, 10, 40, 3, 40, 3, 40, 3, 40, 2, 3, 56, 41, 2,
	4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
	42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76,
	78, 2, 10, 3, 2, 13, 14, 3, 2, 15, 20, 4, 2, 45, 46, 57, 57, 3, 2, 47,
	49, 3, 2, 45, 46, 4, 2, 50, 53, 55, 56, 3, 2, 30, 31, 3, 2, 41, 43, 2,
	455, 2, 86, 3, 2, 2, 2, 4, 91, 3, 2, 2, 2, 6, 102, 3, 2, 2, 2, 8, 111,
	3, 2, 2, 2, 10, 114, 3, 2, 2, 2, 12, 127, 3, 2, 2, 2, 14, 129, 3, 2, 2,
	2, 16, 132, 3, 2, 2, 2, 18, 140, 3, 2, 2, 2, 20, 148, 3, 2, 2, 2, 22, 150,
	3, 2, 2, 2, 24, 163, 3, 2, 2, 2, 26, 203, 3, 2, 2, 2, 28, 205, 3, 2, 2,
	2, 30, 212, 3, 2, 2, 2, 32, 217, 3, 2, 2, 2, 34, 219, 3, 2, 2, 2, 36, 228,
	3, 2, 2, 2, 38, 230, 3, 2, 2, 2, 40, 237, 3, 2, 2, 2, 42, 244, 3, 2, 2,
	2, 44, 256, 3, 2, 2, 2, 46, 258, 3, 2, 2, 2, 48, 260, 3, 2, 2, 2, 50, 266,
	3, 2, 2, 2, 52, 269, 3, 2, 2, 2, 54, 272, 3, 2, 2, 2, 56, 281, 3, 2, 2,
	2, 58, 364, 3, 2, 2, 2, 60, 372, 3, 2, 2, 2, 62, 374, 3, 2, 2, 2, 64, 376,
	3, 2, 2, 2, 66, 389, 3, 2, 2, 2, 68, 391, 3, 2, 2, 2, 70, 394, 3, 2, 2,
	2, 72, 406, 3, 2, 2, 2, 74, 408, 3, 2, 2, 2, 76, 412, 3, 2, 2, 2, 78, 419,
	3, 2, 2, 2, 80, 85, 5, 4, 3, 2, 81, 82, 5, 14, 8, 2, 82, 83, 7, 3, 2, 2,
	83, 85, 3, 2, 2, 2, 84, 80, 3, 2, 2, 2, 84, 81, 3, 2, 2, 2, 85, 88, 3,
	2, 2, 2, 86, 84, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 89, 3, 2, 2, 2, 88,
	86, 3, 2, 2, 2, 89, 90, 7, 2, 2, 3, 90, 3, 3, 2, 2, 2, 91, 92, 7, 4, 2,
	2, 92, 93, 7, 38, 2, 2, 93, 97, 5, 6, 4, 2, 94, 96, 5, 26, 14, 2, 95, 94,
	3, 2, 2, 2, 96, 99, 3, 2, 2, 2, 97, 95, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2,
	98, 100, 3, 2, 2, 2, 99, 97, 3, 2, 2, 2, 100, 101, 5, 10, 6, 2, 101, 5,
	3, 2, 2, 2, 102, 106, 7, 5, 2, 2, 103, 105, 5, 8, 5, 2, 104, 103, 3, 2,
	2, 2, 105, 108, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2,
	107, 109, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 109, 110, 7, 6, 2, 2, 110,
	7, 3, 2, 2, 2, 111, 112, 5, 26, 14, 2, 112, 113, 7, 38, 2, 2, 113, 9, 3,
	2, 2, 2, 114, 118, 7, 7, 2, 2, 115, 117, 5, 12, 7, 2, 116, 115, 3, 2, 2,
	2, 117, 120, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119,
	121, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 121, 122, 7, 8, 2, 2, 122, 11, 3,
	2, 2, 2, 123, 124, 5, 14, 8, 2, 124, 125, 7, 3, 2, 2, 125, 128, 3, 2, 2,
	2, 126, 128, 5, 36, 19, 2, 127, 123, 3, 2, 2, 2, 127, 126, 3, 2, 2, 2,
	128, 13, 3, 2, 2, 2, 129, 130, 5, 26, 14, 2, 130, 131, 5, 16, 9, 2, 131,
	15, 3, 2, 2, 2, 132, 137, 5, 18, 10, 2, 133, 134, 7, 9, 2, 2, 134, 136,
	5, 18, 10, 2, 135, 133, 3, 2, 2, 2, 136, 139, 3, 2, 2, 2, 137, 135, 3,
	2, 2, 2, 137, 138, 3, 2, 2, 2, 138, 17, 3, 2, 2, 2, 139, 137, 3, 2, 2,
	2, 140, 143, 7, 38, 2, 2, 141, 142, 7, 54, 2, 2, 142, 144, 5, 20, 11, 2,
	143, 141, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 19, 3, 2, 2, 2, 145, 149,
	5, 22, 12, 2, 146, 149, 5, 24, 13, 2, 147, 149, 5, 56, 29, 2, 148, 145,
	3, 2, 2, 2, 148, 146, 3, 2, 2, 2, 148, 147, 3, 2, 2, 2, 149, 21, 3, 2,
	2, 2, 150, 159, 7, 7, 2, 2, 151, 156, 5, 20, 11, 2, 152, 153, 7, 9, 2,
	2, 153, 155, 5, 20, 11, 2, 154, 152, 3, 2, 2, 2, 155, 158, 3, 2, 2, 2,
	156, 154, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 160, 3, 2, 2, 2, 158,
	156, 3, 2, 2, 2, 159, 151, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160, 161,
	3, 2, 2, 2, 161, 162, 7, 8, 2, 2, 162, 23, 3, 2, 2, 2, 163, 170, 7, 7,
	2, 2, 164, 165, 5, 56, 29, 2, 165, 166, 7, 10, 2, 2, 166, 167, 5, 20, 11,
	2, 167, 169, 3, 2, 2, 2, 168, 164, 3, 2, 2, 2, 169, 172, 3, 2, 2, 2, 170,
	168, 3, 2, 2, 2, 170, 171, 3, 2, 2, 2, 171, 173, 3, 2, 2, 2, 172, 170,
	3, 2, 2, 2, 173, 174, 7, 8, 2, 2, 174, 25, 3, 2, 2, 2, 175, 179, 5, 34,
	18, 2, 176, 178, 7, 39, 2, 2, 177, 176, 3, 2, 2, 2, 178, 181, 3, 2, 2,
	2, 179, 177, 3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180, 204, 3, 2, 2, 2, 181,
	179, 3, 2, 2, 2, 182, 186, 5, 28, 15, 2, 183, 185, 7, 39, 2, 2, 184, 183,
	3, 2, 2, 2, 185, 188, 3, 2, 2, 2, 186, 184, 3, 2, 2, 2, 186, 187, 3, 2,
	2, 2, 187, 204, 3, 2, 2, 2, 188, 186, 3, 2, 2, 2, 189, 193, 5, 30, 16,
	2, 190, 192, 7, 39, 2, 2, 191, 190, 3, 2, 2, 2, 192, 195, 3, 2, 2, 2, 193,
	191, 3, 2, 2, 2, 193, 194, 3, 2, 2, 2, 194, 204, 3, 2, 2, 2, 195, 193,
	3, 2, 2, 2, 196, 200, 5, 32, 17, 2, 197, 199, 7, 39, 2, 2, 198, 197, 3,
	2, 2, 2, 199, 202, 3, 2, 2, 2, 200, 198, 3, 2, 2, 2, 200, 201, 3, 2, 2,
	2, 201, 204, 3, 2, 2, 2, 202, 200, 3, 2, 2, 2, 203, 175, 3, 2, 2, 2, 203,
	182, 3, 2, 2, 2, 203, 189, 3, 2, 2, 2, 203, 196, 3, 2, 2, 2, 204, 27, 3,
	2, 2, 2, 205, 206, 7, 11, 2, 2, 206, 207, 7, 52, 2, 2, 207, 208, 5, 34,
	18, 2, 208, 209, 7, 9, 2, 2, 209, 210, 5, 26, 14, 2, 210, 211, 7, 50, 2,
	2, 211, 29, 3, 2, 2, 2, 212, 213, 7, 12, 2, 2, 213, 214, 7, 52, 2, 2, 214,
	215, 7, 38, 2, 2, 215, 216, 7, 50, 2, 2, 216, 31, 3, 2, 2, 2, 217, 218,
	9, 2, 2, 2, 218, 33, 3, 2, 2, 2, 219, 220, 9, 3, 2, 2, 220, 35, 3, 2, 2,
	2, 221, 229, 5, 10, 6, 2, 222, 229, 5, 38, 20, 2, 223, 229, 5, 40, 21,
	2, 224, 229, 5, 48, 25, 2, 225, 229, 5, 50, 26, 2, 226, 229, 5, 52, 27,
	2, 227, 229, 5, 54, 28, 2, 228, 221, 3, 2, 2, 2, 228, 222, 3, 2, 2, 2,
	228, 223, 3, 2, 2, 2, 228, 224, 3, 2, 2, 2, 228, 225, 3, 2, 2, 2, 228,
	226, 3, 2, 2, 2, 228, 227, 3, 2, 2, 2, 229, 37, 3, 2, 2, 2, 230, 231, 7,
	21, 2, 2, 231, 232, 5, 56, 29, 2, 232, 235, 5, 36, 19, 2, 233, 234, 7,
	22, 2, 2, 234, 236, 5, 36, 19, 2, 235, 233, 3, 2, 2, 2, 235, 236, 3, 2,
	2, 2, 236, 39, 3, 2, 2, 2, 237, 238, 7, 23, 2, 2, 238, 239, 7, 5, 2, 2,
	239, 240, 5, 42, 22, 2, 240, 241, 7, 6, 2, 2, 241, 242, 5, 36, 19, 2, 242,
	41, 3, 2, 2, 2, 243, 245, 5, 44, 23, 2, 244, 243, 3, 2, 2, 2, 244, 245,
	3, 2, 2, 2, 245, 246, 3, 2, 2, 2, 246, 248, 7, 3, 2, 2, 247, 249, 5, 56,
	29, 2, 248, 247, 3, 2, 2, 2, 248, 249, 3, 2, 2, 2, 249, 250, 3, 2, 2, 2,
	250, 252, 7, 3, 2, 2, 251, 253, 5, 46, 24, 2, 252, 251, 3, 2, 2, 2, 252,
	253, 3, 2, 2, 2, 253, 43, 3, 2, 2, 2, 254, 257, 5, 14, 8, 2, 255, 257,
	5, 64, 33, 2, 256, 254, 3, 2, 2, 2, 256, 255, 3, 2, 2, 2, 257, 45, 3, 2,
	2, 2, 258, 259, 5, 64, 33, 2, 259, 47, 3, 2, 2, 2, 260, 262, 7, 24, 2,
	2, 261, 263, 5, 56, 29, 2, 262, 261, 3, 2, 2, 2, 262, 263, 3, 2, 2, 2,
	263, 264, 3, 2, 2, 2, 264, 265, 7, 3, 2, 2, 265, 49, 3, 2, 2, 2, 266, 267,
	7, 25, 2, 2, 267, 268, 7, 3, 2, 2, 268, 51, 3, 2, 2, 2, 269, 270, 7, 26,
	2, 2, 270, 271, 7, 3, 2, 2, 271, 53, 3, 2, 2, 2, 272, 273, 5, 56, 29, 2,
	273, 274, 7, 3, 2, 2, 274, 55, 3, 2, 2, 2, 275, 276, 8, 29, 1, 2, 276,
	282, 5, 58, 30, 2, 277, 278, 9, 4, 2, 2, 278, 282, 5, 56, 29, 12, 279,
	280, 7, 32, 2, 2, 280, 282, 5, 66, 34, 2, 281, 275, 3, 2, 2, 2, 281, 277,
	3, 2, 2, 2, 281, 279, 3, 2, 2, 2, 282, 355, 3, 2, 2, 2, 283, 284, 12, 16,
	2, 2, 284, 285, 7, 27, 2, 2, 285, 354, 5, 56, 29, 17, 286, 287, 12, 10,
	2, 2, 287, 288, 9, 5, 2, 2, 288, 354, 5, 56, 29, 11, 289, 290, 12, 9, 2,
	2, 290, 291, 9, 6, 2, 2, 291, 354, 5, 56, 29, 10, 292, 293, 12, 8, 2, 2,
	293, 294, 9, 7, 2, 2, 294, 354, 5, 56, 29, 9, 295, 296, 12, 7, 2, 2, 296,
	297, 7, 33, 2, 2, 297, 354, 5, 56, 29, 8, 298, 299, 12, 6, 2, 2, 299, 300,
	7, 34, 2, 2, 300, 354, 5, 56, 29, 7, 301, 302, 12, 5, 2, 2, 302, 303, 7,
	35, 2, 2, 303, 304, 5, 56, 29, 2, 304, 305, 7, 10, 2, 2, 305, 306, 5, 56,
	29, 6, 306, 354, 3, 2, 2, 2, 307, 308, 12, 15, 2, 2, 308, 309, 7, 28, 2,
	2, 309, 310, 5, 56, 29, 2, 310, 311, 7, 29, 2, 2, 311, 354, 3, 2, 2, 2,
	312, 313, 12, 14, 2, 2, 313, 314, 7, 5, 2, 2, 314, 315, 5, 64, 33, 2, 315,
	316, 7, 6, 2, 2, 316, 354, 3, 2, 2, 2, 317, 318, 12, 13, 2, 2, 318, 354,
	9, 8, 2, 2, 319, 324, 12, 4, 2, 2, 320, 321, 7, 9, 2, 2, 321, 323, 5, 56,
	29, 2, 322, 320, 3, 2, 2, 2, 323, 326, 3, 2, 2, 2, 324, 322, 3, 2, 2, 2,
	324, 325, 3, 2, 2, 2, 325, 327, 3, 2, 2, 2, 326, 324, 3, 2, 2, 2, 327,
	328, 7, 54, 2, 2, 328, 333, 5, 56, 29, 2, 329, 330, 7, 9, 2, 2, 330, 332,
	5, 56, 29, 2, 331, 329, 3, 2, 2, 2, 332, 335, 3, 2, 2, 2, 333, 331, 3,
	2, 2, 2, 333, 334, 3, 2, 2, 2, 334, 354, 3, 2, 2, 2, 335, 333, 3, 2, 2,
	2, 336, 341, 12, 3, 2, 2, 337, 338, 7, 9, 2, 2, 338, 340, 5, 56, 29, 2,
	339, 337, 3, 2, 2, 2, 340, 343, 3, 2, 2, 2, 341, 339, 3, 2, 2, 2, 341,
	342, 3, 2, 2, 2, 342, 344, 3, 2, 2, 2, 343, 341, 3, 2, 2, 2, 344, 345,
	7, 36, 2, 2, 345, 350, 5, 56, 29, 2, 346, 347, 7, 9, 2, 2, 347, 349, 5,
	56, 29, 2, 348, 346, 3, 2, 2, 2, 349, 352, 3, 2, 2, 2, 350, 348, 3, 2,
	2, 2, 350, 351, 3, 2, 2, 2, 351, 354, 3, 2, 2, 2, 352, 350, 3, 2, 2, 2,
	353, 283, 3, 2, 2, 2, 353, 286, 3, 2, 2, 2, 353, 289, 3, 2, 2, 2, 353,
	292, 3, 2, 2, 2, 353, 295, 3, 2, 2, 2, 353, 298, 3, 2, 2, 2, 353, 301,
	3, 2, 2, 2, 353, 307, 3, 2, 2, 2, 353, 312, 3, 2, 2, 2, 353, 317, 3, 2,
	2, 2, 353, 319, 3, 2, 2, 2, 353, 336, 3, 2, 2, 2, 354, 357, 3, 2, 2, 2,
	355, 353, 3, 2, 2, 2, 355, 356, 3, 2, 2, 2, 356, 57, 3, 2, 2, 2, 357, 355,
	3, 2, 2, 2, 358, 359, 7, 5, 2, 2, 359, 360, 5, 56, 29, 2, 360, 361, 7,
	6, 2, 2, 361, 365, 3, 2, 2, 2, 362, 365, 5, 60, 31, 2, 363, 365, 7, 38,
	2, 2, 364, 358, 3, 2, 2, 2, 364, 362, 3, 2, 2, 2, 364, 363, 3, 2, 2, 2,
	365, 59, 3, 2, 2, 2, 366, 373, 5, 62, 32, 2, 367, 373, 7, 44, 2, 2, 368,
	373, 7, 58, 2, 2, 369, 373, 7, 59, 2, 2, 370, 373, 7, 37, 2, 2, 371, 373,
	7, 40, 2, 2, 372, 366, 3, 2, 2, 2, 372, 367, 3, 2, 2, 2, 372, 368, 3, 2,
	2, 2, 372, 369, 3, 2, 2, 2, 372, 370, 3, 2, 2, 2, 372, 371, 3, 2, 2, 2,
	373, 61, 3, 2, 2, 2, 374, 375, 9, 9, 2, 2, 375, 63, 3, 2, 2, 2, 376, 381,
	5, 56, 29, 2, 377, 378, 7, 9, 2, 2, 378, 380, 5, 56, 29, 2, 379, 377, 3,
	2, 2, 2, 380, 383, 3, 2, 2, 2, 381, 379, 3, 2, 2, 2, 381, 382, 3, 2, 2,
	2, 382, 65, 3, 2, 2, 2, 383, 381, 3, 2, 2, 2, 384, 390, 5, 68, 35, 2, 385,
	390, 5, 70, 36, 2, 386, 390, 5, 74, 38, 2, 387, 390, 5, 76, 39, 2, 388,
	390, 5, 78, 40, 2, 389, 384, 3, 2, 2, 2, 389, 385, 3, 2, 2, 2, 389, 386,
	3, 2, 2, 2, 389, 387, 3, 2, 2, 2, 389, 388, 3, 2, 2, 2, 390, 67, 3, 2,
	2, 2, 391, 392, 5, 28, 15, 2, 392, 393, 5, 24, 13, 2, 393, 69, 3, 2, 2,
	2, 394, 396, 5, 72, 37, 2, 395, 397, 7, 39, 2, 2, 396, 395, 3, 2, 2, 2,
	397, 398, 3, 2, 2, 2, 398, 396, 3, 2, 2, 2, 398, 399, 3, 2, 2, 2, 399,
	400, 3, 2, 2, 2, 400, 401, 5, 22, 12, 2, 401, 71, 3, 2, 2, 2, 402, 407,
	5, 34, 18, 2, 403, 407, 5, 28, 15, 2, 404, 407, 5, 30, 16, 2, 405, 407,
	5, 32, 17, 2, 406, 402, 3, 2, 2, 2, 406, 403, 3, 2, 2, 2, 406, 404, 3,
	2, 2, 2, 406, 405, 3, 2, 2, 2, 407, 73, 3, 2, 2, 2, 408, 409, 5, 30, 16,
	2, 409, 410, 7, 5, 2, 2, 410, 411, 7, 6, 2, 2, 411, 75, 3, 2, 2, 2, 412,
	413, 5, 34, 18, 2, 413, 415, 7, 5, 2, 2, 414, 416, 5, 56, 29, 2, 415, 414,
	3, 2, 2, 2, 415, 416, 3, 2, 2, 2, 416, 417, 3, 2, 2, 2, 417, 418, 7, 6,
	2, 2, 418, 77, 3, 2, 2, 2, 419, 420, 5, 32, 17, 2, 420, 422, 7, 5, 2, 2,
	421, 423, 5, 56, 29, 2, 422, 421, 3, 2, 2, 2, 422, 423, 3, 2, 2, 2, 423,
	424, 3, 2, 2, 2, 424, 425, 7, 6, 2, 2, 425, 79, 3, 2, 2, 2, 41, 84, 86,
	97, 106, 118, 127, 137, 143, 148, 156, 159, 170, 179, 186, 193, 200, 203,
	228, 235, 244, 248, 252, 256, 262, 281, 324, 333, 341, 350, 353, 355, 364,
	372, 381, 389, 398, 406, 415, 422,
}
var literalNames = []string{
	"", "';'", "'func'", "'('", "')'", "'{'", "'}'", "','", "':'", "'map'",
	"'connector'", "'dynamic'", "'error'", "'int'", "'uint'", "'float'", "'bool'",
	"'char'", "'string'", "'if'", "'else'", "'for'", "'return'", "'break'",
	"'continue'", "'.'", "'['", "']'", "'++'", "'--'", "'new'", "'&&'", "'||'",
	"'?'", "':='", "", "", "", "'null'", "", "", "", "", "'+'", "'-'", "'*'",
	"'/'", "'%'", "'>'", "'>='", "'<'", "'<='", "'='", "'=='", "'!='", "'!'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "BooleanLiteral",
	"Identifier", "Brackets", "NULL", "HexLiteral", "DecimalLiteral", "OctalLiteral",
	"FloatingPointLiteral", "Plus", "Minus", "Multiplication", "Division",
	"Percent", "Gt", "Get", "Lt", "Let", "Assign", "Eq", "NEq", "Not", "CharacterLiteral",
	"StringLiteral", "COMMENT", "WS", "LINE_COMMENT",
}

var ruleNames = []string{
	"compilationUnit", "functionDeclaration", "formalParameters", "formalParameterDecl",
	"block", "blockStatement", "variableDeclaration", "variableDeclarators",
	"variableDeclarator", "variableInitializer", "arrayInitializer", "mapInitializer",
	"type_", "mapType", "connectorType", "dynamicType", "primitiveType", "statement",
	"ifStatement", "forStatement", "forControl", "forInit", "forUpdate", "returnStatement",
	"breakStatement", "continueStatement", "expressionStatement", "expression",
	"primary", "literal", "integerLiteral", "expressionList", "creator", "mapCreator",
	"arrayCreator", "creatorName", "connectorCreator", "primitiveCreator",
	"dynamicCreator",
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
	GoScriptParserT__30                = 31
	GoScriptParserT__31                = 32
	GoScriptParserT__32                = 33
	GoScriptParserT__33                = 34
	GoScriptParserBooleanLiteral       = 35
	GoScriptParserIdentifier           = 36
	GoScriptParserBrackets             = 37
	GoScriptParserNULL                 = 38
	GoScriptParserHexLiteral           = 39
	GoScriptParserDecimalLiteral       = 40
	GoScriptParserOctalLiteral         = 41
	GoScriptParserFloatingPointLiteral = 42
	GoScriptParserPlus                 = 43
	GoScriptParserMinus                = 44
	GoScriptParserMultiplication       = 45
	GoScriptParserDivision             = 46
	GoScriptParserPercent              = 47
	GoScriptParserGt                   = 48
	GoScriptParserGet                  = 49
	GoScriptParserLt                   = 50
	GoScriptParserLet                  = 51
	GoScriptParserAssign               = 52
	GoScriptParserEq                   = 53
	GoScriptParserNEq                  = 54
	GoScriptParserNot                  = 55
	GoScriptParserCharacterLiteral     = 56
	GoScriptParserStringLiteral        = 57
	GoScriptParserCOMMENT              = 58
	GoScriptParserWS                   = 59
	GoScriptParserLINE_COMMENT         = 60
)

// GoScriptParser rules.
const (
	GoScriptParserRULE_compilationUnit     = 0
	GoScriptParserRULE_functionDeclaration = 1
	GoScriptParserRULE_formalParameters    = 2
	GoScriptParserRULE_formalParameterDecl = 3
	GoScriptParserRULE_block               = 4
	GoScriptParserRULE_blockStatement      = 5
	GoScriptParserRULE_variableDeclaration = 6
	GoScriptParserRULE_variableDeclarators = 7
	GoScriptParserRULE_variableDeclarator  = 8
	GoScriptParserRULE_variableInitializer = 9
	GoScriptParserRULE_arrayInitializer    = 10
	GoScriptParserRULE_mapInitializer      = 11
	GoScriptParserRULE_type_               = 12
	GoScriptParserRULE_mapType             = 13
	GoScriptParserRULE_connectorType       = 14
	GoScriptParserRULE_dynamicType         = 15
	GoScriptParserRULE_primitiveType       = 16
	GoScriptParserRULE_statement           = 17
	GoScriptParserRULE_ifStatement         = 18
	GoScriptParserRULE_forStatement        = 19
	GoScriptParserRULE_forControl          = 20
	GoScriptParserRULE_forInit             = 21
	GoScriptParserRULE_forUpdate           = 22
	GoScriptParserRULE_returnStatement     = 23
	GoScriptParserRULE_breakStatement      = 24
	GoScriptParserRULE_continueStatement   = 25
	GoScriptParserRULE_expressionStatement = 26
	GoScriptParserRULE_expression          = 27
	GoScriptParserRULE_primary             = 28
	GoScriptParserRULE_literal             = 29
	GoScriptParserRULE_integerLiteral      = 30
	GoScriptParserRULE_expressionList      = 31
	GoScriptParserRULE_creator             = 32
	GoScriptParserRULE_mapCreator          = 33
	GoScriptParserRULE_arrayCreator        = 34
	GoScriptParserRULE_creatorName         = 35
	GoScriptParserRULE_connectorCreator    = 36
	GoScriptParserRULE_primitiveCreator    = 37
	GoScriptParserRULE_dynamicCreator      = 38
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
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GoScriptParserT__1)|(1<<GoScriptParserT__8)|(1<<GoScriptParserT__9)|(1<<GoScriptParserT__10)|(1<<GoScriptParserT__11)|(1<<GoScriptParserT__12)|(1<<GoScriptParserT__13)|(1<<GoScriptParserT__14)|(1<<GoScriptParserT__15)|(1<<GoScriptParserT__16)|(1<<GoScriptParserT__17))) != 0 {
		p.SetState(82)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case GoScriptParserT__1:
			{
				p.SetState(78)
				p.FunctionDeclaration()
			}

		case GoScriptParserT__8, GoScriptParserT__9, GoScriptParserT__10, GoScriptParserT__11, GoScriptParserT__12, GoScriptParserT__13, GoScriptParserT__14, GoScriptParserT__15, GoScriptParserT__16, GoScriptParserT__17:
			{
				p.SetState(79)
				p.VariableDeclaration()
			}
			{
				p.SetState(80)
				p.Match(GoScriptParserT__0)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(87)
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

func (s *FunctionDeclarationContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionDeclarationContext) AllType_() []IType_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IType_Context)(nil)).Elem())
	var tst = make([]IType_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IType_Context)
		}
	}

	return tst
}

func (s *FunctionDeclarationContext) Type_(i int) IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
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
		p.SetState(89)
		p.Match(GoScriptParserT__1)
	}
	{
		p.SetState(90)
		p.Match(GoScriptParserIdentifier)
	}
	{
		p.SetState(91)
		p.FormalParameters()
	}
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GoScriptParserT__8)|(1<<GoScriptParserT__9)|(1<<GoScriptParserT__10)|(1<<GoScriptParserT__11)|(1<<GoScriptParserT__12)|(1<<GoScriptParserT__13)|(1<<GoScriptParserT__14)|(1<<GoScriptParserT__15)|(1<<GoScriptParserT__16)|(1<<GoScriptParserT__17))) != 0 {
		{
			p.SetState(92)
			p.Type_()
		}

		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(98)
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
		p.SetState(100)
		p.Match(GoScriptParserT__2)
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GoScriptParserT__8)|(1<<GoScriptParserT__9)|(1<<GoScriptParserT__10)|(1<<GoScriptParserT__11)|(1<<GoScriptParserT__12)|(1<<GoScriptParserT__13)|(1<<GoScriptParserT__14)|(1<<GoScriptParserT__15)|(1<<GoScriptParserT__16)|(1<<GoScriptParserT__17))) != 0 {
		{
			p.SetState(101)
			p.FormalParameterDecl()
		}

		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(107)
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
		p.SetState(109)
		p.Type_()
	}
	{
		p.SetState(110)
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
		p.SetState(112)
		p.Match(GoScriptParserT__4)
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GoScriptParserT__2)|(1<<GoScriptParserT__4)|(1<<GoScriptParserT__8)|(1<<GoScriptParserT__9)|(1<<GoScriptParserT__10)|(1<<GoScriptParserT__11)|(1<<GoScriptParserT__12)|(1<<GoScriptParserT__13)|(1<<GoScriptParserT__14)|(1<<GoScriptParserT__15)|(1<<GoScriptParserT__16)|(1<<GoScriptParserT__17)|(1<<GoScriptParserT__18)|(1<<GoScriptParserT__20)|(1<<GoScriptParserT__21)|(1<<GoScriptParserT__22)|(1<<GoScriptParserT__23)|(1<<GoScriptParserT__29))) != 0) || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(GoScriptParserBooleanLiteral-35))|(1<<(GoScriptParserIdentifier-35))|(1<<(GoScriptParserNULL-35))|(1<<(GoScriptParserHexLiteral-35))|(1<<(GoScriptParserDecimalLiteral-35))|(1<<(GoScriptParserOctalLiteral-35))|(1<<(GoScriptParserFloatingPointLiteral-35))|(1<<(GoScriptParserPlus-35))|(1<<(GoScriptParserMinus-35))|(1<<(GoScriptParserNot-35))|(1<<(GoScriptParserCharacterLiteral-35))|(1<<(GoScriptParserStringLiteral-35)))) != 0) {
		{
			p.SetState(113)
			p.BlockStatement()
		}

		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(119)
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

	p.SetState(125)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoScriptParserT__8, GoScriptParserT__9, GoScriptParserT__10, GoScriptParserT__11, GoScriptParserT__12, GoScriptParserT__13, GoScriptParserT__14, GoScriptParserT__15, GoScriptParserT__16, GoScriptParserT__17:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.VariableDeclaration()
		}
		{
			p.SetState(122)
			p.Match(GoScriptParserT__0)
		}

	case GoScriptParserT__2, GoScriptParserT__4, GoScriptParserT__18, GoScriptParserT__20, GoScriptParserT__21, GoScriptParserT__22, GoScriptParserT__23, GoScriptParserT__29, GoScriptParserBooleanLiteral, GoScriptParserIdentifier, GoScriptParserNULL, GoScriptParserHexLiteral, GoScriptParserDecimalLiteral, GoScriptParserOctalLiteral, GoScriptParserFloatingPointLiteral, GoScriptParserPlus, GoScriptParserMinus, GoScriptParserNot, GoScriptParserCharacterLiteral, GoScriptParserStringLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(124)
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
		p.SetState(127)
		p.Type_()
	}
	{
		p.SetState(128)
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
		p.SetState(130)
		p.VariableDeclarator()
	}
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoScriptParserT__6 {
		{
			p.SetState(131)
			p.Match(GoScriptParserT__6)
		}
		{
			p.SetState(132)
			p.VariableDeclarator()
		}

		p.SetState(137)
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
		p.SetState(138)
		p.Match(GoScriptParserIdentifier)
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GoScriptParserAssign {
		{
			p.SetState(139)
			p.Match(GoScriptParserAssign)
		}
		{
			p.SetState(140)
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

	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(143)
			p.ArrayInitializer()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(144)
			p.MapInitializer()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(145)
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

func (s *ArrayInitializerContext) AllVariableInitializer() []IVariableInitializerContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableInitializerContext)(nil)).Elem())
	var tst = make([]IVariableInitializerContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableInitializerContext)
		}
	}

	return tst
}

func (s *ArrayInitializerContext) VariableInitializer(i int) IVariableInitializerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableInitializerContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableInitializerContext)
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
		p.SetState(148)
		p.Match(GoScriptParserT__4)
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GoScriptParserT__2)|(1<<GoScriptParserT__4)|(1<<GoScriptParserT__29))) != 0) || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(GoScriptParserBooleanLiteral-35))|(1<<(GoScriptParserIdentifier-35))|(1<<(GoScriptParserNULL-35))|(1<<(GoScriptParserHexLiteral-35))|(1<<(GoScriptParserDecimalLiteral-35))|(1<<(GoScriptParserOctalLiteral-35))|(1<<(GoScriptParserFloatingPointLiteral-35))|(1<<(GoScriptParserPlus-35))|(1<<(GoScriptParserMinus-35))|(1<<(GoScriptParserNot-35))|(1<<(GoScriptParserCharacterLiteral-35))|(1<<(GoScriptParserStringLiteral-35)))) != 0) {
		{
			p.SetState(149)
			p.VariableInitializer()
		}
		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GoScriptParserT__6 {
			{
				p.SetState(150)
				p.Match(GoScriptParserT__6)
			}
			{
				p.SetState(151)
				p.VariableInitializer()
			}

			p.SetState(156)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(159)
		p.Match(GoScriptParserT__5)
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
	p.EnterRule(localctx, 22, GoScriptParserRULE_mapInitializer)
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
		p.SetState(161)
		p.Match(GoScriptParserT__4)
	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoScriptParserT__2 || _la == GoScriptParserT__29 || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(GoScriptParserBooleanLiteral-35))|(1<<(GoScriptParserIdentifier-35))|(1<<(GoScriptParserNULL-35))|(1<<(GoScriptParserHexLiteral-35))|(1<<(GoScriptParserDecimalLiteral-35))|(1<<(GoScriptParserOctalLiteral-35))|(1<<(GoScriptParserFloatingPointLiteral-35))|(1<<(GoScriptParserPlus-35))|(1<<(GoScriptParserMinus-35))|(1<<(GoScriptParserNot-35))|(1<<(GoScriptParserCharacterLiteral-35))|(1<<(GoScriptParserStringLiteral-35)))) != 0) {
		{
			p.SetState(162)
			p.expression(0)
		}
		{
			p.SetState(163)
			p.Match(GoScriptParserT__7)
		}
		{
			p.SetState(164)
			p.VariableInitializer()
		}

		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(171)
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

func (s *Type_Context) AllBrackets() []antlr.TerminalNode {
	return s.GetTokens(GoScriptParserBrackets)
}

func (s *Type_Context) Brackets(i int) antlr.TerminalNode {
	return s.GetToken(GoScriptParserBrackets, i)
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
	p.EnterRule(localctx, 24, GoScriptParserRULE_type_)
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

	p.SetState(201)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoScriptParserT__12, GoScriptParserT__13, GoScriptParserT__14, GoScriptParserT__15, GoScriptParserT__16, GoScriptParserT__17:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(173)
			p.PrimitiveType()
		}
		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GoScriptParserBrackets {
			{
				p.SetState(174)
				p.Match(GoScriptParserBrackets)
			}

			p.SetState(179)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case GoScriptParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(180)
			p.MapType()
		}
		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GoScriptParserBrackets {
			{
				p.SetState(181)
				p.Match(GoScriptParserBrackets)
			}

			p.SetState(186)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case GoScriptParserT__9:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(187)
			p.ConnectorType()
		}
		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GoScriptParserBrackets {
			{
				p.SetState(188)
				p.Match(GoScriptParserBrackets)
			}

			p.SetState(193)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case GoScriptParserT__10, GoScriptParserT__11:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(194)
			p.DynamicType()
		}
		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GoScriptParserBrackets {
			{
				p.SetState(195)
				p.Match(GoScriptParserBrackets)
			}

			p.SetState(200)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
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

func (s *MapTypeContext) PrimitiveType() IPrimitiveTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitiveTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitiveTypeContext)
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
	p.EnterRule(localctx, 26, GoScriptParserRULE_mapType)

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
		p.SetState(203)
		p.Match(GoScriptParserT__8)
	}
	{
		p.SetState(204)
		p.Match(GoScriptParserLt)
	}
	{
		p.SetState(205)
		p.PrimitiveType()
	}
	{
		p.SetState(206)
		p.Match(GoScriptParserT__6)
	}
	{
		p.SetState(207)
		p.Type_()
	}
	{
		p.SetState(208)
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
	p.EnterRule(localctx, 28, GoScriptParserRULE_connectorType)

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
		p.SetState(210)
		p.Match(GoScriptParserT__9)
	}
	{
		p.SetState(211)
		p.Match(GoScriptParserLt)
	}
	{
		p.SetState(212)
		p.Match(GoScriptParserIdentifier)
	}
	{
		p.SetState(213)
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
	p.EnterRule(localctx, 30, GoScriptParserRULE_dynamicType)
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
		p.SetState(215)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GoScriptParserT__10 || _la == GoScriptParserT__11) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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
	p.EnterRule(localctx, 32, GoScriptParserRULE_primitiveType)
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
		p.SetState(217)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GoScriptParserT__12)|(1<<GoScriptParserT__13)|(1<<GoScriptParserT__14)|(1<<GoScriptParserT__15)|(1<<GoScriptParserT__16)|(1<<GoScriptParserT__17))) != 0) {
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

func (s *StatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) ForStatement() IForStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForStatementContext)
}

func (s *StatementContext) ReturnStatement() IReturnStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *StatementContext) BreakStatement() IBreakStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBreakStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBreakStatementContext)
}

func (s *StatementContext) ContinueStatement() IContinueStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContinueStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContinueStatementContext)
}

func (s *StatementContext) ExpressionStatement() IExpressionStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GoScriptParserRULE_statement)

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

	p.SetState(226)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoScriptParserT__4:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(219)
			p.Block()
		}

	case GoScriptParserT__18:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(220)
			p.IfStatement()
		}

	case GoScriptParserT__20:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(221)
			p.ForStatement()
		}

	case GoScriptParserT__21:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(222)
			p.ReturnStatement()
		}

	case GoScriptParserT__22:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(223)
			p.BreakStatement()
		}

	case GoScriptParserT__23:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(224)
			p.ContinueStatement()
		}

	case GoScriptParserT__2, GoScriptParserT__29, GoScriptParserBooleanLiteral, GoScriptParserIdentifier, GoScriptParserNULL, GoScriptParserHexLiteral, GoScriptParserDecimalLiteral, GoScriptParserOctalLiteral, GoScriptParserFloatingPointLiteral, GoScriptParserPlus, GoScriptParserMinus, GoScriptParserNot, GoScriptParserCharacterLiteral, GoScriptParserStringLiteral:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(225)
			p.ExpressionStatement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_ifStatement
	return p
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

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

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GoScriptParserRULE_ifStatement)

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
		p.SetState(228)
		p.Match(GoScriptParserT__18)
	}
	{
		p.SetState(229)
		p.expression(0)
	}
	{
		p.SetState(230)
		p.Statement()
	}
	p.SetState(233)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(231)
			p.Match(GoScriptParserT__19)
		}
		{
			p.SetState(232)
			p.Statement()
		}

	}

	return localctx
}

// IForStatementContext is an interface to support dynamic dispatch.
type IForStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForStatementContext differentiates from other interfaces.
	IsForStatementContext()
}

type ForStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStatementContext() *ForStatementContext {
	var p = new(ForStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_forStatement
	return p
}

func (*ForStatementContext) IsForStatementContext() {}

func NewForStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStatementContext {
	var p = new(ForStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_forStatement

	return p
}

func (s *ForStatementContext) GetParser() antlr.Parser { return s.parser }

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

func (s *ForStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitForStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) ForStatement() (localctx IForStatementContext) {
	localctx = NewForStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GoScriptParserRULE_forStatement)

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
		p.SetState(235)
		p.Match(GoScriptParserT__20)
	}
	{
		p.SetState(236)
		p.Match(GoScriptParserT__2)
	}
	{
		p.SetState(237)
		p.ForControl()
	}
	{
		p.SetState(238)
		p.Match(GoScriptParserT__3)
	}
	{
		p.SetState(239)
		p.Statement()
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
	p.EnterRule(localctx, 40, GoScriptParserRULE_forControl)
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
	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GoScriptParserT__2)|(1<<GoScriptParserT__8)|(1<<GoScriptParserT__9)|(1<<GoScriptParserT__10)|(1<<GoScriptParserT__11)|(1<<GoScriptParserT__12)|(1<<GoScriptParserT__13)|(1<<GoScriptParserT__14)|(1<<GoScriptParserT__15)|(1<<GoScriptParserT__16)|(1<<GoScriptParserT__17)|(1<<GoScriptParserT__29))) != 0) || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(GoScriptParserBooleanLiteral-35))|(1<<(GoScriptParserIdentifier-35))|(1<<(GoScriptParserNULL-35))|(1<<(GoScriptParserHexLiteral-35))|(1<<(GoScriptParserDecimalLiteral-35))|(1<<(GoScriptParserOctalLiteral-35))|(1<<(GoScriptParserFloatingPointLiteral-35))|(1<<(GoScriptParserPlus-35))|(1<<(GoScriptParserMinus-35))|(1<<(GoScriptParserNot-35))|(1<<(GoScriptParserCharacterLiteral-35))|(1<<(GoScriptParserStringLiteral-35)))) != 0) {
		{
			p.SetState(241)
			p.ForInit()
		}

	}
	{
		p.SetState(244)
		p.Match(GoScriptParserT__0)
	}
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GoScriptParserT__2 || _la == GoScriptParserT__29 || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(GoScriptParserBooleanLiteral-35))|(1<<(GoScriptParserIdentifier-35))|(1<<(GoScriptParserNULL-35))|(1<<(GoScriptParserHexLiteral-35))|(1<<(GoScriptParserDecimalLiteral-35))|(1<<(GoScriptParserOctalLiteral-35))|(1<<(GoScriptParserFloatingPointLiteral-35))|(1<<(GoScriptParserPlus-35))|(1<<(GoScriptParserMinus-35))|(1<<(GoScriptParserNot-35))|(1<<(GoScriptParserCharacterLiteral-35))|(1<<(GoScriptParserStringLiteral-35)))) != 0) {
		{
			p.SetState(245)
			p.expression(0)
		}

	}
	{
		p.SetState(248)
		p.Match(GoScriptParserT__0)
	}
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GoScriptParserT__2 || _la == GoScriptParserT__29 || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(GoScriptParserBooleanLiteral-35))|(1<<(GoScriptParserIdentifier-35))|(1<<(GoScriptParserNULL-35))|(1<<(GoScriptParserHexLiteral-35))|(1<<(GoScriptParserDecimalLiteral-35))|(1<<(GoScriptParserOctalLiteral-35))|(1<<(GoScriptParserFloatingPointLiteral-35))|(1<<(GoScriptParserPlus-35))|(1<<(GoScriptParserMinus-35))|(1<<(GoScriptParserNot-35))|(1<<(GoScriptParserCharacterLiteral-35))|(1<<(GoScriptParserStringLiteral-35)))) != 0) {
		{
			p.SetState(249)
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
	p.EnterRule(localctx, 42, GoScriptParserRULE_forInit)

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

	p.SetState(254)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoScriptParserT__8, GoScriptParserT__9, GoScriptParserT__10, GoScriptParserT__11, GoScriptParserT__12, GoScriptParserT__13, GoScriptParserT__14, GoScriptParserT__15, GoScriptParserT__16, GoScriptParserT__17:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(252)
			p.VariableDeclaration()
		}

	case GoScriptParserT__2, GoScriptParserT__29, GoScriptParserBooleanLiteral, GoScriptParserIdentifier, GoScriptParserNULL, GoScriptParserHexLiteral, GoScriptParserDecimalLiteral, GoScriptParserOctalLiteral, GoScriptParserFloatingPointLiteral, GoScriptParserPlus, GoScriptParserMinus, GoScriptParserNot, GoScriptParserCharacterLiteral, GoScriptParserStringLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(253)
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
	p.EnterRule(localctx, 44, GoScriptParserRULE_forUpdate)

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
		p.SetState(256)
		p.ExpressionList()
	}

	return localctx
}

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_returnStatement
	return p
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, GoScriptParserRULE_returnStatement)
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
		p.SetState(258)
		p.Match(GoScriptParserT__21)
	}
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GoScriptParserT__2 || _la == GoScriptParserT__29 || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(GoScriptParserBooleanLiteral-35))|(1<<(GoScriptParserIdentifier-35))|(1<<(GoScriptParserNULL-35))|(1<<(GoScriptParserHexLiteral-35))|(1<<(GoScriptParserDecimalLiteral-35))|(1<<(GoScriptParserOctalLiteral-35))|(1<<(GoScriptParserFloatingPointLiteral-35))|(1<<(GoScriptParserPlus-35))|(1<<(GoScriptParserMinus-35))|(1<<(GoScriptParserNot-35))|(1<<(GoScriptParserCharacterLiteral-35))|(1<<(GoScriptParserStringLiteral-35)))) != 0) {
		{
			p.SetState(259)
			p.expression(0)
		}

	}
	{
		p.SetState(262)
		p.Match(GoScriptParserT__0)
	}

	return localctx
}

// IBreakStatementContext is an interface to support dynamic dispatch.
type IBreakStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBreakStatementContext differentiates from other interfaces.
	IsBreakStatementContext()
}

type BreakStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakStatementContext() *BreakStatementContext {
	var p = new(BreakStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_breakStatement
	return p
}

func (*BreakStatementContext) IsBreakStatementContext() {}

func NewBreakStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStatementContext {
	var p = new(BreakStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_breakStatement

	return p
}

func (s *BreakStatementContext) GetParser() antlr.Parser { return s.parser }
func (s *BreakStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitBreakStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) BreakStatement() (localctx IBreakStatementContext) {
	localctx = NewBreakStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, GoScriptParserRULE_breakStatement)

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
		p.SetState(264)
		p.Match(GoScriptParserT__22)
	}
	{
		p.SetState(265)
		p.Match(GoScriptParserT__0)
	}

	return localctx
}

// IContinueStatementContext is an interface to support dynamic dispatch.
type IContinueStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContinueStatementContext differentiates from other interfaces.
	IsContinueStatementContext()
}

type ContinueStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinueStatementContext() *ContinueStatementContext {
	var p = new(ContinueStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_continueStatement
	return p
}

func (*ContinueStatementContext) IsContinueStatementContext() {}

func NewContinueStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinueStatementContext {
	var p = new(ContinueStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_continueStatement

	return p
}

func (s *ContinueStatementContext) GetParser() antlr.Parser { return s.parser }
func (s *ContinueStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinueStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitContinueStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) ContinueStatement() (localctx IContinueStatementContext) {
	localctx = NewContinueStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, GoScriptParserRULE_continueStatement)

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
		p.SetState(267)
		p.Match(GoScriptParserT__23)
	}
	{
		p.SetState(268)
		p.Match(GoScriptParserT__0)
	}

	return localctx
}

// IExpressionStatementContext is an interface to support dynamic dispatch.
type IExpressionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionStatementContext differentiates from other interfaces.
	IsExpressionStatementContext()
}

type ExpressionStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionStatementContext() *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_expressionStatement
	return p
}

func (*ExpressionStatementContext) IsExpressionStatementContext() {}

func NewExpressionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_expressionStatement

	return p
}

func (s *ExpressionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitExpressionStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) ExpressionStatement() (localctx IExpressionStatementContext) {
	localctx = NewExpressionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, GoScriptParserRULE_expressionStatement)

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
		p.SetState(270)
		p.expression(0)
	}
	{
		p.SetState(271)
		p.Match(GoScriptParserT__0)
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

type MulExprContext struct {
	*ExpressionContext
}

func NewMulExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulExprContext {
	var p = new(MulExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MulExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MulExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MulExprContext) Multiplication() antlr.TerminalNode {
	return s.GetToken(GoScriptParserMultiplication, 0)
}

func (s *MulExprContext) Division() antlr.TerminalNode {
	return s.GetToken(GoScriptParserDivision, 0)
}

func (s *MulExprContext) Percent() antlr.TerminalNode {
	return s.GetToken(GoScriptParserPercent, 0)
}

func (s *MulExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitMulExpr(s)

	default:
		return t.VisitChildren(s)
	}
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

type CreateAndAssignExprContext struct {
	*ExpressionContext
}

func NewCreateAndAssignExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CreateAndAssignExprContext {
	var p = new(CreateAndAssignExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *CreateAndAssignExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateAndAssignExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *CreateAndAssignExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CreateAndAssignExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitCreateAndAssignExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddExprContext struct {
	*ExpressionContext
}

func NewAddExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddExprContext {
	var p = new(AddExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AddExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AddExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddExprContext) Plus() antlr.TerminalNode {
	return s.GetToken(GoScriptParserPlus, 0)
}

func (s *AddExprContext) Minus() antlr.TerminalNode {
	return s.GetToken(GoScriptParserMinus, 0)
}

func (s *AddExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitAddExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConditionalExprContext struct {
	*ExpressionContext
}

func NewConditionalExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConditionalExprContext {
	var p = new(ConditionalExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ConditionalExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ConditionalExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionalExprContext) Let() antlr.TerminalNode {
	return s.GetToken(GoScriptParserLet, 0)
}

func (s *ConditionalExprContext) Get() antlr.TerminalNode {
	return s.GetToken(GoScriptParserGet, 0)
}

func (s *ConditionalExprContext) Gt() antlr.TerminalNode {
	return s.GetToken(GoScriptParserGt, 0)
}

func (s *ConditionalExprContext) Lt() antlr.TerminalNode {
	return s.GetToken(GoScriptParserLt, 0)
}

func (s *ConditionalExprContext) Eq() antlr.TerminalNode {
	return s.GetToken(GoScriptParserEq, 0)
}

func (s *ConditionalExprContext) NEq() antlr.TerminalNode {
	return s.GetToken(GoScriptParserNEq, 0)
}

func (s *ConditionalExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitConditionalExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryExprContext struct {
	*ExpressionContext
}

func NewUnaryExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryExprContext {
	var p = new(UnaryExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExprContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryExprContext) Plus() antlr.TerminalNode {
	return s.GetToken(GoScriptParserPlus, 0)
}

func (s *UnaryExprContext) Minus() antlr.TerminalNode {
	return s.GetToken(GoScriptParserMinus, 0)
}

func (s *UnaryExprContext) Not() antlr.TerminalNode {
	return s.GetToken(GoScriptParserNot, 0)
}

func (s *UnaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitUnaryExpr(s)

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

type SelectorExprContext struct {
	*ExpressionContext
}

func NewSelectorExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectorExprContext {
	var p = new(SelectorExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *SelectorExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorExprContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *SelectorExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SelectorExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitSelectorExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CreateExprContext struct {
	*ExpressionContext
}

func NewCreateExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CreateExprContext {
	var p = new(CreateExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *CreateExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateExprContext) Creator() ICreatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICreatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICreatorContext)
}

func (s *CreateExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitCreateExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type SelfAddExprContext struct {
	*ExpressionContext
}

func NewSelfAddExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelfAddExprContext {
	var p = new(SelfAddExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *SelfAddExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelfAddExprContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SelfAddExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitSelfAddExpr(s)

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
	_startState := 54
	p.EnterRecursionRule(localctx, 54, GoScriptParserRULE_expression, _p)
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
	p.SetState(279)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoScriptParserT__2, GoScriptParserBooleanLiteral, GoScriptParserIdentifier, GoScriptParserNULL, GoScriptParserHexLiteral, GoScriptParserDecimalLiteral, GoScriptParserOctalLiteral, GoScriptParserFloatingPointLiteral, GoScriptParserCharacterLiteral, GoScriptParserStringLiteral:
		localctx = NewPrimaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(274)
			p.Primary()
		}

	case GoScriptParserPlus, GoScriptParserMinus, GoScriptParserNot:
		localctx = NewUnaryExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(275)
			_la = p.GetTokenStream().LA(1)

			if !(((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(GoScriptParserPlus-43))|(1<<(GoScriptParserMinus-43))|(1<<(GoScriptParserNot-43)))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(276)
			p.expression(10)
		}

	case GoScriptParserT__29:
		localctx = NewCreateExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(277)
			p.Match(GoScriptParserT__29)
		}
		{
			p.SetState(278)
			p.Creator()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(351)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
			case 1:
				localctx = NewSelectorExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(281)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(282)
					p.Match(GoScriptParserT__24)
				}
				{
					p.SetState(283)
					p.expression(15)
				}

			case 2:
				localctx = NewMulExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(284)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(285)
					_la = p.GetTokenStream().LA(1)

					if !(((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(GoScriptParserMultiplication-45))|(1<<(GoScriptParserDivision-45))|(1<<(GoScriptParserPercent-45)))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(286)
					p.expression(9)
				}

			case 3:
				localctx = NewAddExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(287)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(288)
					_la = p.GetTokenStream().LA(1)

					if !(_la == GoScriptParserPlus || _la == GoScriptParserMinus) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(289)
					p.expression(8)
				}

			case 4:
				localctx = NewConditionalExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(290)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(291)
					_la = p.GetTokenStream().LA(1)

					if !(((_la-48)&-(0x1f+1)) == 0 && ((1<<uint((_la-48)))&((1<<(GoScriptParserGt-48))|(1<<(GoScriptParserGet-48))|(1<<(GoScriptParserLt-48))|(1<<(GoScriptParserLet-48))|(1<<(GoScriptParserEq-48))|(1<<(GoScriptParserNEq-48)))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(292)
					p.expression(7)
				}

			case 5:
				localctx = NewAndExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(293)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(294)
					p.Match(GoScriptParserT__30)
				}
				{
					p.SetState(295)
					p.expression(6)
				}

			case 6:
				localctx = NewOrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(296)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(297)
					p.Match(GoScriptParserT__31)
				}
				{
					p.SetState(298)
					p.expression(5)
				}

			case 7:
				localctx = NewTernaryExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(299)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(300)
					p.Match(GoScriptParserT__32)
				}
				{
					p.SetState(301)
					p.expression(0)
				}
				{
					p.SetState(302)
					p.Match(GoScriptParserT__7)
				}
				{
					p.SetState(303)
					p.expression(4)
				}

			case 8:
				localctx = NewIndexExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(305)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(306)
					p.Match(GoScriptParserT__25)
				}
				{
					p.SetState(307)
					p.expression(0)
				}
				{
					p.SetState(308)
					p.Match(GoScriptParserT__26)
				}

			case 9:
				localctx = NewCallExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(310)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(311)
					p.Match(GoScriptParserT__2)
				}
				{
					p.SetState(312)
					p.ExpressionList()
				}
				{
					p.SetState(313)
					p.Match(GoScriptParserT__3)
				}

			case 10:
				localctx = NewSelfAddExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(315)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(316)
					_la = p.GetTokenStream().LA(1)

					if !(_la == GoScriptParserT__27 || _la == GoScriptParserT__28) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			case 11:
				localctx = NewAssignExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(317)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				p.SetState(322)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for _la == GoScriptParserT__6 {
					{
						p.SetState(318)
						p.Match(GoScriptParserT__6)
					}
					{
						p.SetState(319)
						p.expression(0)
					}

					p.SetState(324)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(325)
					p.Match(GoScriptParserAssign)
				}
				{
					p.SetState(326)
					p.expression(0)
				}
				p.SetState(331)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())

				for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					if _alt == 1 {
						{
							p.SetState(327)
							p.Match(GoScriptParserT__6)
						}
						{
							p.SetState(328)
							p.expression(0)
						}

					}
					p.SetState(333)
					p.GetErrorHandler().Sync(p)
					_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())
				}

			case 12:
				localctx = NewCreateAndAssignExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GoScriptParserRULE_expression)
				p.SetState(334)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				p.SetState(339)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for _la == GoScriptParserT__6 {
					{
						p.SetState(335)
						p.Match(GoScriptParserT__6)
					}
					{
						p.SetState(336)
						p.expression(0)
					}

					p.SetState(341)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(342)
					p.Match(GoScriptParserT__33)
				}
				{
					p.SetState(343)
					p.expression(0)
				}
				p.SetState(348)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())

				for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					if _alt == 1 {
						{
							p.SetState(344)
							p.Match(GoScriptParserT__6)
						}
						{
							p.SetState(345)
							p.expression(0)
						}

					}
					p.SetState(350)
					p.GetErrorHandler().Sync(p)
					_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())
				}

			}

		}
		p.SetState(355)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 56, GoScriptParserRULE_primary)

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

	p.SetState(362)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoScriptParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(356)
			p.Match(GoScriptParserT__2)
		}
		{
			p.SetState(357)
			p.expression(0)
		}
		{
			p.SetState(358)
			p.Match(GoScriptParserT__3)
		}

	case GoScriptParserBooleanLiteral, GoScriptParserNULL, GoScriptParserHexLiteral, GoScriptParserDecimalLiteral, GoScriptParserOctalLiteral, GoScriptParserFloatingPointLiteral, GoScriptParserCharacterLiteral, GoScriptParserStringLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(360)
			p.Literal()
		}

	case GoScriptParserIdentifier:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(361)
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
	p.EnterRule(localctx, 58, GoScriptParserRULE_literal)

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

	p.SetState(370)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoScriptParserHexLiteral, GoScriptParserDecimalLiteral, GoScriptParserOctalLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(364)
			p.IntegerLiteral()
		}

	case GoScriptParserFloatingPointLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(365)
			p.Match(GoScriptParserFloatingPointLiteral)
		}

	case GoScriptParserCharacterLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(366)
			p.Match(GoScriptParserCharacterLiteral)
		}

	case GoScriptParserStringLiteral:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(367)
			p.Match(GoScriptParserStringLiteral)
		}

	case GoScriptParserBooleanLiteral:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(368)
			p.Match(GoScriptParserBooleanLiteral)
		}

	case GoScriptParserNULL:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(369)
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
	p.EnterRule(localctx, 60, GoScriptParserRULE_integerLiteral)
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
		p.SetState(372)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-39)&-(0x1f+1)) == 0 && ((1<<uint((_la-39)))&((1<<(GoScriptParserHexLiteral-39))|(1<<(GoScriptParserDecimalLiteral-39))|(1<<(GoScriptParserOctalLiteral-39)))) != 0) {
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
	p.EnterRule(localctx, 62, GoScriptParserRULE_expressionList)
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
		p.SetState(374)
		p.expression(0)
	}
	p.SetState(379)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GoScriptParserT__6 {
		{
			p.SetState(375)
			p.Match(GoScriptParserT__6)
		}
		{
			p.SetState(376)
			p.expression(0)
		}

		p.SetState(381)
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

func (s *CreatorContext) MapCreator() IMapCreatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapCreatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapCreatorContext)
}

func (s *CreatorContext) ArrayCreator() IArrayCreatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayCreatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayCreatorContext)
}

func (s *CreatorContext) ConnectorCreator() IConnectorCreatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConnectorCreatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConnectorCreatorContext)
}

func (s *CreatorContext) PrimitiveCreator() IPrimitiveCreatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitiveCreatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitiveCreatorContext)
}

func (s *CreatorContext) DynamicCreator() IDynamicCreatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDynamicCreatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDynamicCreatorContext)
}

func (s *CreatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitCreator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) Creator() (localctx ICreatorContext) {
	localctx = NewCreatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, GoScriptParserRULE_creator)

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

	p.SetState(387)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(382)
			p.MapCreator()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(383)
			p.ArrayCreator()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(384)
			p.ConnectorCreator()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(385)
			p.PrimitiveCreator()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(386)
			p.DynamicCreator()
		}

	}

	return localctx
}

// IMapCreatorContext is an interface to support dynamic dispatch.
type IMapCreatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapCreatorContext differentiates from other interfaces.
	IsMapCreatorContext()
}

type MapCreatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapCreatorContext() *MapCreatorContext {
	var p = new(MapCreatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_mapCreator
	return p
}

func (*MapCreatorContext) IsMapCreatorContext() {}

func NewMapCreatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapCreatorContext {
	var p = new(MapCreatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_mapCreator

	return p
}

func (s *MapCreatorContext) GetParser() antlr.Parser { return s.parser }

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

func (s *MapCreatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapCreatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapCreatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitMapCreator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) MapCreator() (localctx IMapCreatorContext) {
	localctx = NewMapCreatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, GoScriptParserRULE_mapCreator)

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
		p.SetState(389)
		p.MapType()
	}
	{
		p.SetState(390)
		p.MapInitializer()
	}

	return localctx
}

// IArrayCreatorContext is an interface to support dynamic dispatch.
type IArrayCreatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayCreatorContext differentiates from other interfaces.
	IsArrayCreatorContext()
}

type ArrayCreatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayCreatorContext() *ArrayCreatorContext {
	var p = new(ArrayCreatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_arrayCreator
	return p
}

func (*ArrayCreatorContext) IsArrayCreatorContext() {}

func NewArrayCreatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayCreatorContext {
	var p = new(ArrayCreatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_arrayCreator

	return p
}

func (s *ArrayCreatorContext) GetParser() antlr.Parser { return s.parser }

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

func (s *ArrayCreatorContext) AllBrackets() []antlr.TerminalNode {
	return s.GetTokens(GoScriptParserBrackets)
}

func (s *ArrayCreatorContext) Brackets(i int) antlr.TerminalNode {
	return s.GetToken(GoScriptParserBrackets, i)
}

func (s *ArrayCreatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayCreatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayCreatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitArrayCreator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) ArrayCreator() (localctx IArrayCreatorContext) {
	localctx = NewArrayCreatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, GoScriptParserRULE_arrayCreator)
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
		p.SetState(392)
		p.CreatorName()
	}
	p.SetState(394)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GoScriptParserBrackets {
		{
			p.SetState(393)
			p.Match(GoScriptParserBrackets)
		}

		p.SetState(396)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(398)
		p.ArrayInitializer()
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
	p.EnterRule(localctx, 70, GoScriptParserRULE_creatorName)

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

	p.SetState(404)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GoScriptParserT__12, GoScriptParserT__13, GoScriptParserT__14, GoScriptParserT__15, GoScriptParserT__16, GoScriptParserT__17:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(400)
			p.PrimitiveType()
		}

	case GoScriptParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(401)
			p.MapType()
		}

	case GoScriptParserT__9:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(402)
			p.ConnectorType()
		}

	case GoScriptParserT__10, GoScriptParserT__11:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(403)
			p.DynamicType()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConnectorCreatorContext is an interface to support dynamic dispatch.
type IConnectorCreatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConnectorCreatorContext differentiates from other interfaces.
	IsConnectorCreatorContext()
}

type ConnectorCreatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConnectorCreatorContext() *ConnectorCreatorContext {
	var p = new(ConnectorCreatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_connectorCreator
	return p
}

func (*ConnectorCreatorContext) IsConnectorCreatorContext() {}

func NewConnectorCreatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConnectorCreatorContext {
	var p = new(ConnectorCreatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_connectorCreator

	return p
}

func (s *ConnectorCreatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ConnectorCreatorContext) ConnectorType() IConnectorTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConnectorTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConnectorTypeContext)
}

func (s *ConnectorCreatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConnectorCreatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConnectorCreatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitConnectorCreator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) ConnectorCreator() (localctx IConnectorCreatorContext) {
	localctx = NewConnectorCreatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, GoScriptParserRULE_connectorCreator)

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
		p.SetState(406)
		p.ConnectorType()
	}
	{
		p.SetState(407)
		p.Match(GoScriptParserT__2)
	}
	{
		p.SetState(408)
		p.Match(GoScriptParserT__3)
	}

	return localctx
}

// IPrimitiveCreatorContext is an interface to support dynamic dispatch.
type IPrimitiveCreatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimitiveCreatorContext differentiates from other interfaces.
	IsPrimitiveCreatorContext()
}

type PrimitiveCreatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitiveCreatorContext() *PrimitiveCreatorContext {
	var p = new(PrimitiveCreatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_primitiveCreator
	return p
}

func (*PrimitiveCreatorContext) IsPrimitiveCreatorContext() {}

func NewPrimitiveCreatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveCreatorContext {
	var p = new(PrimitiveCreatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_primitiveCreator

	return p
}

func (s *PrimitiveCreatorContext) GetParser() antlr.Parser { return s.parser }

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

func (s *PrimitiveCreatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveCreatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitiveCreatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitPrimitiveCreator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) PrimitiveCreator() (localctx IPrimitiveCreatorContext) {
	localctx = NewPrimitiveCreatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, GoScriptParserRULE_primitiveCreator)
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
		p.SetState(410)
		p.PrimitiveType()
	}
	{
		p.SetState(411)
		p.Match(GoScriptParserT__2)
	}
	p.SetState(413)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GoScriptParserT__2 || _la == GoScriptParserT__29 || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(GoScriptParserBooleanLiteral-35))|(1<<(GoScriptParserIdentifier-35))|(1<<(GoScriptParserNULL-35))|(1<<(GoScriptParserHexLiteral-35))|(1<<(GoScriptParserDecimalLiteral-35))|(1<<(GoScriptParserOctalLiteral-35))|(1<<(GoScriptParserFloatingPointLiteral-35))|(1<<(GoScriptParserPlus-35))|(1<<(GoScriptParserMinus-35))|(1<<(GoScriptParserNot-35))|(1<<(GoScriptParserCharacterLiteral-35))|(1<<(GoScriptParserStringLiteral-35)))) != 0) {
		{
			p.SetState(412)
			p.expression(0)
		}

	}
	{
		p.SetState(415)
		p.Match(GoScriptParserT__3)
	}

	return localctx
}

// IDynamicCreatorContext is an interface to support dynamic dispatch.
type IDynamicCreatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDynamicCreatorContext differentiates from other interfaces.
	IsDynamicCreatorContext()
}

type DynamicCreatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDynamicCreatorContext() *DynamicCreatorContext {
	var p = new(DynamicCreatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GoScriptParserRULE_dynamicCreator
	return p
}

func (*DynamicCreatorContext) IsDynamicCreatorContext() {}

func NewDynamicCreatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DynamicCreatorContext {
	var p = new(DynamicCreatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoScriptParserRULE_dynamicCreator

	return p
}

func (s *DynamicCreatorContext) GetParser() antlr.Parser { return s.parser }

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

func (s *DynamicCreatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DynamicCreatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DynamicCreatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoScriptVisitor:
		return t.VisitDynamicCreator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoScriptParser) DynamicCreator() (localctx IDynamicCreatorContext) {
	localctx = NewDynamicCreatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, GoScriptParserRULE_dynamicCreator)
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
		p.SetState(417)
		p.DynamicType()
	}
	{
		p.SetState(418)
		p.Match(GoScriptParserT__2)
	}
	p.SetState(420)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GoScriptParserT__2 || _la == GoScriptParserT__29 || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(GoScriptParserBooleanLiteral-35))|(1<<(GoScriptParserIdentifier-35))|(1<<(GoScriptParserNULL-35))|(1<<(GoScriptParserHexLiteral-35))|(1<<(GoScriptParserDecimalLiteral-35))|(1<<(GoScriptParserOctalLiteral-35))|(1<<(GoScriptParserFloatingPointLiteral-35))|(1<<(GoScriptParserPlus-35))|(1<<(GoScriptParserMinus-35))|(1<<(GoScriptParserNot-35))|(1<<(GoScriptParserCharacterLiteral-35))|(1<<(GoScriptParserStringLiteral-35)))) != 0) {
		{
			p.SetState(419)
			p.expression(0)
		}

	}
	{
		p.SetState(422)
		p.Match(GoScriptParserT__3)
	}

	return localctx
}

func (p *GoScriptParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 27:
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
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
