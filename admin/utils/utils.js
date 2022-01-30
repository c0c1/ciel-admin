export function bgGround(value, background) {
    if (background) {
        return background
    }
    switch (value%19) {
        case 1:
            return "#cfc"
        case 2:
            return "#ffcdd2"
        case 3:
            return "#ccf"
        case 4:
            return "#e1bee7"
        case 5:
            return "#ede7f6"
        case -1:
        case 6:
            return "#c5cae9"
        case 7:
            return "#bbdefb"
        case 8:
            return "#81d4fa"
        case 9:
            return "#80deea"
        case 10:
            return "#b2dfdb"
        case 11:
            return "#c8e6c9"
        case 12:
            return "#dcedc8"
        case 13:
            return "#f0f4c3"
        case 14:
            return "#b9f6ca"
        case 15:
            return "#ccff90"
        case 16:
            return "#f4ff81"
        case 17:
            return "#ffccbc"
        case 18:
            return "#bcaaa4"
        case 19:
            return "#e0e0e0"
        case 20:
        default:
            return "#cfd8dc"
    }
}
