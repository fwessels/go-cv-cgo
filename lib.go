package gocv

func Convert(src View, dst View) {
	assert(src.Size().Equals(dst.Size()) && src.format != NONE && dst.format != NONE)

	if src.format == dst.format {
		Copy(src, dst)
		return
	}

	switch src.format {
	case GRAY8:
		switch dst.format {
		case BGRA32:
			GrayToBgra(src, dst)
		case BGR24:
			GrayToBgr(src, dst)
		default:
			panic("0")
		}
	case BGR24:
		switch dst.format {
		case BGRA32:
			BgrToBgra(src, dst)
		case GRAY8:
			BgrToGray(src, dst)
		default:
			panic("0")
		}
	case BGRA32:
		switch dst.format {
		case BGR24:
			BgraToBgr(src, dst)
		case GRAY8:
			BgraToGray(src, dst)
		default:
			panic("0")
		}
	default:
		panic("0")
	}
}
