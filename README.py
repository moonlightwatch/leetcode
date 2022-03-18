import os
from typing import List




def get_all_files() -> List[str]:
    files = []
    for folder in os.listdir("leetcode"):
        if not folder.startswith("L_"):
            continue
        for file in os.listdir(os.path.join("leetcode", folder)):
            if file.startswith("L_") and file.endswith(".go"):
                files.append(os.path.join("leetcode", folder, file))
    return files

def make_content() -> str:
    content = "## Leetcode 做题记录\n\n"
    sum = 0
    easy = 0
    medium = 0
    hard = 0
    for f in get_all_files():
        sum+=1
        if "easy" in f:
            easy+=1
        if "medium" in f:
            medium+=1
        if "hard" in f:
            hard+=1
    content += "### 数量统计\n\n"
    content += f"已做题总数：{sum}  \n"
    content += f"容易（easy）：{easy}  \n"
    content += f"中等（medium）：{medium}  \n"
    content += f"困难（hard）：{hard}  \n"
    return content
        



if __name__ == "__main__":
    with open(file="README.md",mode="w",encoding="utf-8") as f:
        f.write(make_content())
        f.write("\n\n")