import matplotlib.pyplot as plt
import json

TEST_DATA = "4h15000.json"
EXTREMA_DATA = "inf.json"


def show(x, y, hy, ly, ex=None, elasty=None, eprevy=None):
    fig = plt.figure()
    ax = fig.add_subplot(111)
    ax.plot(x, y, label="close")
    ax.plot(x, hy, label="high")
    ax.plot(x, ly, label="low")
    if ex and elasty and eprevy:
        ax.plot(ex, elasty, label="localmax")
        ax.plot(ex, eprevy, label="localmin")
    plt.legend()
    plt.grid(True)
    plt.show()


def slice(dic, start):
    ret = {}
    for k in dic.keys():
        if dic[k] is not None and len(dic[k]) >= start:
            ret[k] = dic[k][start:]
    return ret


def main():
    pass


if __name__ == "__main__":
    with open(TEST_DATA) as f:
        tdata = json.load(f)
        tdata["TIndex"] = [i for i in range(len(tdata["time"]))]

    with open(EXTREMA_DATA) as f:
        edata = json.load(f)

    nt = slice(tdata, 14800)
    ne = slice(edata, 14800)

    # show(nt["time"], nt["c"], nt["h"], nt["l"])
    last = [d["Last"]["Val"] for d in ne["Infs"]]
    prev = [d["Prev"]["Val"] for d in ne["Infs"]]
    show(nt["TIndex"], nt["c"], nt["h"], nt["l"], ne["TIndex"], last, prev)
