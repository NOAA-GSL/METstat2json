# /// script
# requires-python = ">=3.13"
# dependencies = [
#     "os",
#     "pandas",
#     "json",
#     "re"
# ]
# ///

# Author: Lindsay Blank, CIRA
# Date: February 2026
# Contact: Lindsay.Blank@noaa.gov
# Description: This program reads in every MET output file in a directory, determines the datatype of each column within the file, and outputs a type json file.

import os
import pandas as pd
import json
import re


def read_met_file(file_path):
    """This function reads in a MET output file and returns a pandas dataframe.
       It assumes the file is in fixed-width format, which is the default output format for MET.

    Args:
        file_path (string): The path to the MET output file.

    Returns:
        met_df (Pandas dataframe): A dataframe containing the data from the MET output file.
    """
    met_df = pd.read_fwf(file_path)
    return met_df


def determine_linetype(met_df):
    """This function determines the linetype of a MET output file based on the presence of certain columns or values in the dataframe.

    Args:
        met_df (Pandas dataframe): A dataframe containing MET output.

    Returns:
        linetype (string): The linetype of the MET output file.
    """

    filetype_dict = {
        "STAT": [
            "CNT",
            "CTC",
            "CTS",
            "FHO",
            "ISC",
            "MCTC",
            "MCTS",
            "MPR",
            "SEEPS",
            "SEEPS_MPR",
            "NBRCNT",
            "NBRCTC",
            "NBRCTS",
            "GRAD",
            "DMAP",
            "ORANK",
            "PCT",
            "PJC",
            "PRC",
            "PSTD",
            "ECLV",
            "ECNT",
            "RPS",
            "RHIST",
            "PHIST",
            "RELP",
            "SAL1L2",
            "SL1L2",
            "SSVAR",
            "VAL1L2",
            "VL1L2",
            "VCNT",
            "GENMPR",
            "SSIDX",
        ],
        "TCST": ["TCMPR", "TCDIAG", "PROBRIRW"],
    }

    # The order of the if statements matters here. For example, GSI_MPR has all of the MPR columns plus extra, so we need to check for GSI_MPR first.
    if "QC_WGHT" in met_df.columns:
        linetype = "GSI_MPR_CONVENTIONAL"
    elif "SETUP_QC" in met_df.columns and "OBS_WGHT" not in met_df.columns:
        linetype = "GSI_ORANK_CONVENTIONAL"
    elif "PRS_MAX_WGT" in met_df.columns:
        linetype = "GSI_MPR_RADIANCE"
    elif "SUN_ZNTH" in met_df.columns and "PRS_MAX_WGT" not in met_df.columns:
        linetype = "GSI_ORANK_RADIANCE"
    elif "LINE_TYPE" in met_df.columns:
        if met_df["LINE_TYPE"].iloc[0] in filetype_dict["STAT"]:
            linetype = met_df["LINE_TYPE"].iloc[0]
        elif met_df["LINE_TYPE"].iloc[0] in filetype_dict["TCST"]:
            linetype = met_df["LINE_TYPE"].iloc[0]
    elif "TIME_INDEX" in met_df.columns:
        linetype = "MTD_2DSINGLE"
    elif "CDIST_TRAVELLED" in met_df.columns:
        linetype = "MTD_3DSINGLE"
    elif "SPACE_CENTROID_DIST" in met_df.columns:
        linetype = "MTD_3DPAIR"
    elif "OBJECT_ID" in met_df.columns:
        linetype = "MODE_OBJ"
    elif "FY_OY" in met_df.columns:
        linetype = "MODE_CTS"
    else:
        linetype = "ERROR"
    return linetype


def determine_column_types(met_df):
    """This function determines the datatype of each column in a MET output file.
    It forces certain columns to be specific datatypes based on prior knowledge of the MET output format.

    Args:
        met_df (Pandas dataframe): A dataframe containing MET output.

    Returns:
        column_types (dictionary): A dictionary where the keys are column names and the values are the datatypes of those columns.
    """
    # Columns we know need to be a specific datatype.
    force_float_types = ["CSI_NCL", "PODY_NCL", "OBS_ELV", "ENS_MEAN", "ALON", "BLON", "X_ERR"]
    force_string_types = [
        "DESC",
        "FCST_LVL",
        "OBS_LVL",
        "COV_THRESH",
        "OBS_THRESH",
        "FCST_THRESH",
        "FCST_UNITS",
        "OBS_UNITS",
        "FCST_LEV",
        "OBS_LEV",
        "OBTYPE",
        "VX_MASK",
        "INTERP_MTHD",
        "FCST",
        "OBS",
        "OBS_QC",
        "INITIALS"
    ]
    force_int_types = [
        "INTERP_PNTS",
        "OBS_SID",
        "FCST_LEAD",
        "OBS_LEAD",
        "RANK",
        "AREA",
    ]

    column_types = {}
    for column in met_df.columns:
        column_types[column] = str(met_df[column].dtype)
        if column in force_float_types:
            column_types[column] = "float64"
        elif column in force_string_types:
            column_types[column] = "str"
        elif column in force_int_types:
            column_types[column] = "int64"
        elif re.search(r"ENS_\d+", column) or re.search(r"THRESH_\d+", column):
            column_types[column] = "float64"
    return column_types


def are_column_types_consistent(column_types):
    """This function sees if columns with the same name across different MET output files have the same datatype.

    Args:
        column_types (dictionary): A dictionary where the keys are column names and the values are the datatypes of those columns.

    Returns:
        consistent (Boolean): True if columns with the same name across different MET output files have the same datatype.
        unique_types (dictionary): A dictionary where the keys are column names and the values are the unique datatypes of those columns across all MET output files.
    """
    unique_types = {}
    for column_type in column_types:
        for col, dtype in column_type.items():
            if col not in unique_types:
                unique_types[col] = set()
            unique_types[col].add(dtype)

    consistent = all(len(types) == 1 for types in unique_types.values())
    return consistent, unique_types


def main():
    directory = "met_test_files"
    linetypes_dict = {}
    column_types_summary = []
    for filename in os.listdir(directory):
        file_path = os.path.join(directory, filename)
        print(f"Processing file: {filename}")
        met_df = read_met_file(file_path)
        column_types = determine_column_types(met_df)
        column_types_summary.append(column_types)
        linetype = determine_linetype(met_df)
        linetypes_dict[linetype] = {col: dtype for col, dtype in column_types.items()}
        if linetype == "TCMPR":
            linetypes_dict["TCDIAG"] = {col: dtype for col, dtype in column_types.items()}

    consistent, unique_types = are_column_types_consistent(column_types_summary)
    print("Are the column types consistent?: ", consistent)
    print(
        "Datatypes for each column across files with more than one datatype: ",
        {col: types for col, types in unique_types.items() if len(types) > 1},
    )

    # Sort only the linetype keys
    sorted_linetypes_dict = {
        k: linetypes_dict[k] for k in sorted(linetypes_dict.keys())
    }

    output_filename = "met_column_types_v12.2.json"
    with open(output_filename, "w") as f:
        json.dump(sorted_linetypes_dict, f, indent=4)


if __name__ == "__main__":
    main()
