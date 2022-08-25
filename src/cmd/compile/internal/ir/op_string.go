// Code generated by "stringer -type=Op -trimprefix=O node.go"; DO NOT EDIT.

package ir

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OXXX-0]
	_ = x[ONAME-1]
	_ = x[ONONAME-2]
	_ = x[OTYPE-3]
	_ = x[OLITERAL-4]
	_ = x[ONIL-5]
	_ = x[OADD-6]
	_ = x[OSUB-7]
	_ = x[OOR-8]
	_ = x[OXOR-9]
	_ = x[OADDSTR-10]
	_ = x[OADDR-11]
	_ = x[OANDAND-12]
	_ = x[OAPPEND-13]
	_ = x[OBYTES2STR-14]
	_ = x[OBYTES2STRTMP-15]
	_ = x[ORUNES2STR-16]
	_ = x[OSTR2BYTES-17]
	_ = x[OSTR2BYTESTMP-18]
	_ = x[OSTR2RUNES-19]
	_ = x[OSLICE2ARRPTR-20]
	_ = x[OAS-21]
	_ = x[OAS2-22]
	_ = x[OAS2DOTTYPE-23]
	_ = x[OAS2FUNC-24]
	_ = x[OAS2MAPR-25]
	_ = x[OAS2RECV-26]
	_ = x[OASOP-27]
	_ = x[OCALL-28]
	_ = x[OCALLFUNC-29]
	_ = x[OCALLMETH-30]
	_ = x[OCALLINTER-31]
	_ = x[OCAP-32]
	_ = x[OCLOSE-33]
	_ = x[OCLOSURE-34]
	_ = x[OCOMPLIT-35]
	_ = x[OMAPLIT-36]
	_ = x[OSTRUCTLIT-37]
	_ = x[OARRAYLIT-38]
	_ = x[OSLICELIT-39]
	_ = x[OPTRLIT-40]
	_ = x[OCONV-41]
	_ = x[OCONVIFACE-42]
	_ = x[OCONVIDATA-43]
	_ = x[OCONVNOP-44]
	_ = x[OCOPY-45]
	_ = x[ODCL-46]
	_ = x[ODCLFUNC-47]
	_ = x[ODCLCONST-48]
	_ = x[ODCLTYPE-49]
	_ = x[ODELETE-50]
	_ = x[ODOT-51]
	_ = x[ODOTPTR-52]
	_ = x[ODOTMETH-53]
	_ = x[ODOTINTER-54]
	_ = x[OXDOT-55]
	_ = x[ODOTTYPE-56]
	_ = x[ODOTTYPE2-57]
	_ = x[OEQ-58]
	_ = x[ONE-59]
	_ = x[OLT-60]
	_ = x[OLE-61]
	_ = x[OGE-62]
	_ = x[OGT-63]
	_ = x[ODEREF-64]
	_ = x[OINDEX-65]
	_ = x[OINDEXMAP-66]
	_ = x[OKEY-67]
	_ = x[OSTRUCTKEY-68]
	_ = x[OLEN-69]
	_ = x[OMAKE-70]
	_ = x[OMAKECHAN-71]
	_ = x[OMAKEMAP-72]
	_ = x[OMAKESLICE-73]
	_ = x[OMAKESLICECOPY-74]
	_ = x[OMUL-75]
	_ = x[ODIV-76]
	_ = x[OMOD-77]
	_ = x[OLSH-78]
	_ = x[ORSH-79]
	_ = x[OAND-80]
	_ = x[OANDNOT-81]
	_ = x[ONEW-82]
	_ = x[ONOT-83]
	_ = x[OBITNOT-84]
	_ = x[OPLUS-85]
	_ = x[ONEG-86]
	_ = x[OOROR-87]
	_ = x[OPANIC-88]
	_ = x[OPRINT-89]
	_ = x[OPRINTN-90]
	_ = x[OPAREN-91]
	_ = x[OSEND-92]
	_ = x[OSLICE-93]
	_ = x[OSLICEARR-94]
	_ = x[OSLICESTR-95]
	_ = x[OSLICE3-96]
	_ = x[OSLICE3ARR-97]
	_ = x[OSLICEHEADER-98]
	_ = x[ORECOVER-99]
	_ = x[ORECOVERFP-100]
	_ = x[ORECV-101]
	_ = x[ORUNESTR-102]
	_ = x[OSELRECV2-103]
	_ = x[OREAL-104]
	_ = x[OIMAG-105]
	_ = x[OCOMPLEX-106]
	_ = x[OALIGNOF-107]
	_ = x[OOFFSETOF-108]
	_ = x[OSIZEOF-109]
	_ = x[OUNSAFEADD-110]
	_ = x[OUNSAFESLICE-111]
	_ = x[OMETHEXPR-112]
	_ = x[OMETHVALUE-113]
	_ = x[OBLOCK-114]
	_ = x[OBREAK-115]
	_ = x[OCASE-116]
	_ = x[OCONTINUE-117]
	_ = x[ODEFER-118]
	_ = x[OFALL-119]
	_ = x[OFOR-120]
	_ = x[OGOTO-121]
	_ = x[OIF-122]
	_ = x[OLABEL-123]
	_ = x[OGO-124]
	_ = x[ORANGE-125]
	_ = x[ORETURN-126]
	_ = x[OSELECT-127]
	_ = x[OSWITCH-128]
	_ = x[OTYPESW-129]
	_ = x[OFUNCINST-130]
	_ = x[OINLCALL-131]
	_ = x[OEFACE-132]
	_ = x[OITAB-133]
	_ = x[OIDATA-134]
	_ = x[OSPTR-135]
	_ = x[OCFUNC-136]
	_ = x[OCHECKNIL-137]
	_ = x[ORESULT-138]
	_ = x[OINLMARK-139]
	_ = x[OLINKSYMOFFSET-140]
	_ = x[OJUMPTABLE-141]
	_ = x[ODYNAMICDOTTYPE-142]
	_ = x[ODYNAMICDOTTYPE2-143]
	_ = x[ODYNAMICTYPE-144]
	_ = x[OTAILCALL-145]
	_ = x[OGETG-146]
	_ = x[OGETCALLERPC-147]
	_ = x[OGETCALLERSP-148]
	_ = x[OEND-149]
}

const _Op_name = "XXXNAMENONAMETYPELITERALNILADDSUBORXORADDSTRADDRANDANDAPPENDBYTES2STRBYTES2STRTMPRUNES2STRSTR2BYTESSTR2BYTESTMPSTR2RUNESSLICE2ARRPTRASAS2AS2DOTTYPEAS2FUNCAS2MAPRAS2RECVASOPCALLCALLFUNCCALLMETHCALLINTERCAPCLOSECLOSURECOMPLITMAPLITSTRUCTLITARRAYLITSLICELITPTRLITCONVCONVIFACECONVIDATACONVNOPCOPYDCLDCLFUNCDCLCONSTDCLTYPEDELETEDOTDOTPTRDOTMETHDOTINTERXDOTDOTTYPEDOTTYPE2EQNELTLEGEGTDEREFINDEXINDEXMAPKEYSTRUCTKEYLENMAKEMAKECHANMAKEMAPMAKESLICEMAKESLICECOPYMULDIVMODLSHRSHANDANDNOTNEWNOTBITNOTPLUSNEGORORPANICPRINTPRINTNPARENSENDSLICESLICEARRSLICESTRSLICE3SLICE3ARRSLICEHEADERRECOVERRECOVERFPRECVRUNESTRSELRECV2REALIMAGCOMPLEXALIGNOFOFFSETOFSIZEOFUNSAFEADDUNSAFESLICEMETHEXPRMETHVALUEBLOCKBREAKCASECONTINUEDEFERFALLFORGOTOIFLABELGORANGERETURNSELECTSWITCHTYPESWFUNCINSTINLCALLEFACEITABIDATASPTRCFUNCCHECKNILRESULTINLMARKLINKSYMOFFSETJUMPTABLEDYNAMICDOTTYPEDYNAMICDOTTYPE2DYNAMICTYPETAILCALLGETGGETCALLERPCGETCALLERSPEND"

var _Op_index = [...]uint16{0, 3, 7, 13, 17, 24, 27, 30, 33, 35, 38, 44, 48, 54, 60, 69, 81, 90, 99, 111, 120, 132, 134, 137, 147, 154, 161, 168, 172, 176, 184, 192, 201, 204, 209, 216, 223, 229, 238, 246, 254, 260, 264, 273, 282, 289, 293, 296, 303, 311, 318, 324, 327, 333, 340, 348, 352, 359, 367, 369, 371, 373, 375, 377, 379, 384, 389, 397, 400, 409, 412, 416, 424, 431, 440, 453, 456, 459, 462, 465, 468, 471, 477, 480, 483, 489, 493, 496, 500, 505, 510, 516, 521, 525, 530, 538, 546, 552, 561, 572, 579, 588, 592, 599, 607, 611, 615, 622, 629, 637, 643, 652, 663, 671, 680, 685, 690, 694, 702, 707, 711, 714, 718, 720, 725, 727, 732, 738, 744, 750, 756, 764, 771, 776, 780, 785, 789, 794, 802, 808, 815, 828, 837, 851, 866, 877, 885, 889, 900, 911, 914}

func (i Op) String() string {
	if i >= Op(len(_Op_index)-1) {
		return "Op(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Op_name[_Op_index[i]:_Op_index[i+1]]
}
