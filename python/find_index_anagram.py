# Find the indices of all anagrams of a given word in a another word. For
# example: Find the indices of all the anagrams of AB in ABCDBACDAB (Answer: 0, 4, 8)

def solution(word, sample):
    l = len(sample)
    to_return = list()
    for idx in range(len(word) - l + 1):
        w = word[idx:idx+l]
        if anagram(w, sample):
            to_return.append(idx)
    return to_return


def anagram(w1, w2):
    return ''.join(sorted(w1)) == ''.join(sorted(w2))


print(solution("ABCDBACDAB", "AB"))
