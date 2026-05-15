# frozen_string_literal: true

require_relative "junction/version"

module Junction
  class Error < StandardError; end unless const_defined?(:Error)
end
