# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://rubydoc.brew.sh/Formula
# PLEASE REMOVE ALL GENERATED COMMENTS BEFORE SUBMITTING YOUR PR!
class Gmx < Formula
  desc "A Golang based code/file generation tool that feeds data into templates"
  homepage "https://github.com/razpinator/gmx"
  url "https://github.com/razpinator/gmx/archive/v0.0.8.tar.gz"
  sha256 "YOUR_SHA256_HERE"  # You'll need to update this
  license "MIT"  # Update with your actual license

  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args(ldflags: "-s -w")
  end

  test do
    system "#{bin}/gmx", "--help"
  end
end