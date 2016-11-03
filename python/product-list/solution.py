def my_func(input_list):
    output = [1]
    l = len(input_list)
    for idx in range(1, l):
        output.append(output[-1] * input_list[idx-1])

    product = 1
    for idx in range(l-2, -1, -1):
        product *= input_list[idx+1]
        output[idx] *= product

    return output
