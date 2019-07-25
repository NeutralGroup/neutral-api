# -*- coding: utf-8 -*-
#
#   logger
#
# -----------------------------------------------------------------------------

import logging

# -----------------------------------------------------------------------------

# global logging level
logging_level = logging.INFO

def set_logging_level(level):
    logging_level = level

def get_logger(name):
    formatter = logging.Formatter(fmt='%(asctime)s - %(levelname)s - %(module)s - %(message)s')
    handler = logging.StreamHandler()
    handler.setFormatter(formatter)
    logger = logging.getLogger(name)
    logger.setLevel(logging_level)
    logger.addHandler(handler)
    return logger

# =============================================================================
