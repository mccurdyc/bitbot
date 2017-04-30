#' FUNCTION: read_algorithm_data
#'
#' Read the data collected from buying and selling using greedy and random algorithms.
#' @export

read_algorithm_data <- function(d) {
  f <- system.file("extdata", "algorithm-data.csv", package="bitbotRpkg")
  d <- readr::read_csv(f)
  return(dplyr::tbl_df(d))
}
