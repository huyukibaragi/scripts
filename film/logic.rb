require_relative 'common'

def next_page?(doc)
  doc.css('div.p-main-area.p-main-area--wide > div > p').text.include?('一致する情報は見つかりませんでした。')
end

def main(root_url)
  page_number = 1
  charset = nil

  loop do
    url = root_url + "?page=#{page_number}"
    begin
      html = open(url, allow_redirections: :safe) do |f|
        charset = f.charset
        f.read
      end
    rescue StandardError => e
      puts e
      break
    end

    sleep 1 # sleep 1
    doc = Nokogiri::HTML.parse(html, nil, charset)
    break if next_page?(doc)

    doc.css('body > div.l-main > div.p-content.p-content--grid > div.p-main-area.p-main-area--wide > div.p-movies-grid > div').each_with_index do |node, i|
      next if i.zero?
      File.open("url.txt","a") do |text|
        text.puts("https://filmarks.com/movies/#{node.css('div.js-btn-mark.p-movie-cassette__action.p-movie-cassette__action--marks').attribute('data-movie-id').value}")
      end
      #puts node.css('div.p-movie-cassette__info >div.p-movie-cassette__info__main > h3').text
    end
    page_number += 1 # add pageNumber
  end
end
year = 1920
loop do
  break if year == 2030
  puts "https://filmarks.com/list/year/#{year}s"
  main("https://filmarks.com/list/year/#{year}s")
  year += 10
end
