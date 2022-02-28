import argparse

import os
import numpy as np
import onnxruntime as ort
import tokenization
from run_onnx_squad import *
import json

MAX_SEQ_LENGTH = 256
DOC_STRIDE = 128
MAX_QUERY_LENGTH = 64
BATCH_SIZE = 1
N_BEST_SIZE = 20
MAX_ANSWER_LENGTH = 30
BERT_SQUAD_12_ONNX_FILE = "/home/ubuntu/gogogo/python_backend/bertsquad-12.onnx"
VOCAB_FILE = "/home/ubuntu/gogogo/python_backend/vocab.txt"
INPUT_DIR = "/home/ubuntu/server_inputs"
OUTPUT_DIR = "/home/ubuntu/server_outputs"

if __name__ == "__main__":
    parser = argparse.ArgumentParser()

    parser.add_argument(
        "-f",
        "-in_file",
        dest="input_file",
        type=str,
    )
    parser.add_argument(
        "-o",
        "-out_file",
        dest="output_file",
        type=str,
    )

    args = parser.parse_args()
    input_file_name = args.input_file
    output_file_name = args.output_file

    # Use read_squad_examples method from run_onnx_squad to read the input file
    eval_examples = read_squad_examples(
        input_file=os.path.join(INPUT_DIR, input_file_name)
    )
    tokenizer = tokenization.FullTokenizer(vocab_file=VOCAB_FILE, do_lower_case=True)

    # Use convert_examples_to_features method from run_onnx_squad to get parameters from the input
    input_ids, input_mask, segment_ids, extra_data = convert_examples_to_features(
        eval_examples, tokenizer, MAX_SEQ_LENGTH, DOC_STRIDE, MAX_QUERY_LENGTH
    )

    session = ort.InferenceSession(BERT_SQUAD_12_ONNX_FILE)
    n = len(eval_examples)
    bs = BATCH_SIZE
    all_results = []
    for idx in range(0, n):
        item = eval_examples[0]
        # this is using batch_size=1
        # feed the input data as int64
        data = {
            "unique_ids_raw_output___9:0": np.array([item.qas_id], dtype=np.int64),
            "input_ids:0": input_ids[idx : idx + bs],
            "input_mask:0": input_mask[idx : idx + bs],
            "segment_ids:0": segment_ids[idx : idx + bs],
        }
        result = session.run(["unique_ids:0", "unstack:0", "unstack:1"], data)
        in_batch = result[1].shape[0]
        start_logits = [float(x) for x in result[1][0].flat]
        end_logits = [float(x) for x in result[2][0].flat]
        for i in range(0, in_batch):
            unique_id = len(all_results)
            all_results.append(
                RawResult(
                    unique_id=unique_id,
                    start_logits=start_logits,
                    end_logits=end_logits,
                )
            )

    output_prediction_file = os.path.join(OUTPUT_DIR, output_file_name)
    write_predictions(
        eval_examples,
        extra_data,
        all_results,
        N_BEST_SIZE,
        MAX_ANSWER_LENGTH,
        True,
        output_prediction_file,
    )

    with open(output_prediction_file) as json_file:
        answers = json.load(json_file)
        print(json.dumps(answers, indent=2))
